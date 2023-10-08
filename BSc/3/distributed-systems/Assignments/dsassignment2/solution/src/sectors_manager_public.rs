use std::collections::HashMap;
use std::path::PathBuf;
use std::sync::Arc;
use tokio::io::{AsyncReadExt, AsyncWriteExt};
use tokio::sync::RwLock;

use crate::{SectorIdx, SectorVec};

#[async_trait::async_trait]
pub trait SectorsManager: Send + Sync {
    /// Returns 4096 bytes of sector data by index.
    async fn read_data(&self, idx: SectorIdx) -> SectorVec;

    /// Returns timestamp and write rank of the process which has saved this data.
    /// Timestamps and ranks are relevant for atomic register algorithm, and are described
    /// there.
    async fn read_metadata(&self, idx: SectorIdx) -> (u64, u8);

    /// Writes a new data, along with timestamp and write rank to some sector.
    async fn write(&self, idx: SectorIdx, sector: &(SectorVec, u64, u8));
}

/// Path parameter points to a directory to which this method has exclusive access.
pub fn build_sectors_manager(path: PathBuf) -> Arc<dyn SectorsManager> {
    Arc::new(SectorsManagerImpl::new(path))
}

struct SectorsManagerImpl {
    metadata: RwLock<HashMap<SectorIdx, (u64, u8)>>,
    path: PathBuf,
}

impl SectorsManagerImpl {
    pub fn new(path: PathBuf) -> Self {
        let folder_path = path.join(SectorsManagerImpl::tmp_dir());

        let res = Self {
            metadata: RwLock::new(SectorsManagerImpl::restore(&path)),
            path,
        };

        std::fs::create_dir_all(folder_path).unwrap();

        res
    }

    fn tmp_dir() -> &'static str {
        "tmp"
    }

    fn restore(path: &PathBuf) -> HashMap<SectorIdx, (u64, u8)> {
        let map = SectorsManagerImpl::get_hash_map(path);

        SectorsManagerImpl::clean_dir(path, &map);

        map
    }

    fn clean_dir(path: &PathBuf, map: &HashMap<SectorIdx, (u64, u8)>) {
        for entry_ in std::fs::read_dir(path).unwrap() {
            let entry = match entry_ {
                Err(_) => continue,
                Ok(e) => e,
            };

            let path = entry.path();

            let file_name = path.file_name().unwrap().to_string_lossy().into_owned();

            // Tmp dir may have not be destroyed.
            if entry.file_type().unwrap().is_dir() {
                std::fs::remove_dir_all(&path).unwrap();
                continue;
            }

            // Some old state files may have been left.
            let split: Vec<&str> = file_name.split("_").collect();
            if split.len() == 3 {
                let idx = split[0].parse().unwrap();
                let timestamp: u64 = split[1].parse().unwrap();
                let wr: u8 = split[2].parse().unwrap();

                let (newest_ts, newest_wr) = map[&idx];

                if timestamp != newest_ts || wr != newest_wr {
                    std::fs::remove_file(&path).unwrap();
                }
            }
        }
    }

    fn get_hash_map(path: &PathBuf) -> HashMap<SectorIdx, (u64, u8)> {
        let mut res = HashMap::new();

        for entry_ in std::fs::read_dir(path).unwrap() {
            let entry = match entry_ {
                Err(_) => continue,
                Ok(e) => e,
            };

            let file_name = entry
                .path()
                .file_name()
                .unwrap()
                .to_string_lossy()
                .into_owned();

            let split: Vec<&str> = file_name.split("_").collect();
            if split.len() == 3 {
                let idx = split[0].parse().unwrap();
                let timestamp = split[1].parse().unwrap();
                let wr = split[2].parse().unwrap();

                match res.get(&idx) {
                    None => {
                        res.insert(idx, (timestamp, wr));
                    }
                    Some((org_ts, _)) => {
                        if timestamp > *org_ts {
                            res.insert(idx, (timestamp, wr));
                        }
                    }
                }
            }
        }

        res
    }

    fn get_sector_file_name(&self, idx: SectorIdx, timestamp: u64, wr: u8) -> String {
        format!("{}_{}_{}", idx, timestamp, wr)
    }

    fn get_file_path(&self, file_name: &str) -> PathBuf {
        self.path.join(file_name)
    }

    fn get_tmp_path(&self, file_name: &str) -> PathBuf {
        self.path
            .join(SectorsManagerImpl::tmp_dir())
            .join(file_name)
    }

    async fn save_to_file_atomically(
        &self,
        name: &str,
        value: &[u8],
    ) -> Result<(), tokio::io::Error> {
        let tmp_path = self.get_tmp_path(name);

        let mut file = tokio::fs::File::create(&tmp_path).await?;
        file.write_all(value).await?;
        file.sync_data().await?;

        let save_path = self.get_file_path(&name);

        tokio::fs::rename(&tmp_path, &save_path).await?;
        tokio::fs::File::open(save_path)
            .await
            .unwrap()
            .sync_data()
            .await
    }
}

#[async_trait::async_trait]
impl SectorsManager for SectorsManagerImpl {
    /// Returns 4096 bytes of sector data by index.
    async fn read_data(&self, idx: SectorIdx) -> SectorVec {
        match self.metadata.read().await.get(&idx) {
            None => SectorVec(vec![0; 4096]),
            Some((ts, wr)) => {
                let file_name = self.get_sector_file_name(idx, *ts, *wr);

                let mut file = tokio::fs::File::open(self.get_file_path(&file_name))
                    .await
                    .unwrap();

                let mut res = Vec::new();
                res.reserve(4096);

                file.read_to_end(&mut res).await.unwrap();

                SectorVec(res)
            }
        }
    }

    /// Returns timestamp and write rank of the process which has saved this data.
    /// Timestamps and ranks are relevant for atomic register algorithm, and are described
    /// there.
    async fn read_metadata(&self, idx: SectorIdx) -> (u64, u8) {
        match self.metadata.read().await.get(&idx) {
            None => (0, 0),
            Some(res) => *res,
        }
    }

    /// Writes a new data, along with timestamp and write rank to some sector.
    async fn write(&self, idx: SectorIdx, sector: &(SectorVec, u64, u8)) {
        let (SectorVec(vec), ts, wr) = sector;
        let new_name = self.get_sector_file_name(idx, *ts, *wr);

        self.save_to_file_atomically(&new_name, vec).await.unwrap();

        let mut metadata = self.metadata.write().await;
        let prev_val = metadata.insert(idx, (*ts, *wr));
        std::mem::drop(metadata);

        if let Some((prev_ts, prev_wr)) = prev_val {
            let prev_name = self.get_sector_file_name(idx, prev_ts, prev_wr);
            tokio::fs::remove_file(self.get_file_path(&prev_name))
                .await
                .unwrap();
        }
    }
}
