use std::collections::HashMap;
use std::io::{BufRead, BufReader, Read, Write};
use std::path::PathBuf;

use tokio::fs::OpenOptions;
use tokio::io::AsyncWriteExt;

// You can add here other imports from std or crates listed in Cargo.toml.

// You can add any private types, structs, consts, functions, methods, etc., you need.

#[async_trait::async_trait]
pub trait StableStorage: Send + Sync {
    /// Stores `value` under  `key`.
    ///
    /// Detailed requirements are specified in the description of the assignment.
    async fn put(&mut self, key: &str, value: &[u8]) -> Result<(), String>;

    /// Retrieves value stored under `key`.
    ///
    /// Detailed requirements are specified in the description of the assignment.
    async fn get(&self, key: &str) -> Option<Vec<u8>>;
}

struct Storage {
    root_storage_dir: PathBuf,
    name_map: HashMap<String, String>,
}

const KEYS_FILE_NAME: &'static str = "keys";

impl Storage {
    pub fn new(root_storage_dir: PathBuf) -> Self {
        std::fs::create_dir_all(root_storage_dir.join(Storage::tmp_dir())).unwrap();

        Storage {
            root_storage_dir,
            name_map: HashMap::new(),
        }
    }

    pub fn restore(&mut self) {
        let key_file_path = self.get_file_path(KEYS_FILE_NAME);

        match std::fs::File::open(&key_file_path) {
            Err(_) => {
                std::fs::File::create(key_file_path).unwrap();
            }
            Ok(f) => self.restore_map_helper(f),
        }
    }

    fn tmp_dir() -> &'static str {
        "tmp"
    }

    fn restore_map_helper(&mut self, file: std::fs::File) {
        let mut keys_in_order = self.get_map_keys_in_order(file);

        self.clean_map_file(&keys_in_order);

        for (i, key) in keys_in_order.drain(..).enumerate() {
            self.name_map.insert(key, i.to_string());
        }
    }

    fn get_map_keys_in_order(&self, file: std::fs::File) -> Vec<String> {
        let mut reader = BufReader::new(file);

        let mut len_vec = vec![0u8];
        let mut keys = Vec::new();

        while let Ok(_) = reader.read_exact(&mut len_vec) {
            let mut key = String::new();
            match reader.read_line(&mut key) {
                Ok(len) => {
                    if len == 0 {
                        break;
                    }
                    key.pop();
                    if key.len() != len_vec[0] as usize {
                        break;
                    }
                    keys.push(key);
                }
                Err(_) => {
                    break;
                }
            }
        }

        keys
    }

    fn clean_map_file(&self, keys_in_order: &Vec<String>) {
        let tmp_path = self.get_tmp_file_path();
        let mut new_map_file = std::fs::File::create(&tmp_path).unwrap();

        for key in keys_in_order {
            new_map_file
                .write_all(Storage::get_key_line(&key).as_bytes())
                .unwrap();
        }

        new_map_file.sync_data().unwrap();
        let f_path = self.get_file_path(KEYS_FILE_NAME);
        std::fs::rename(tmp_path, &f_path).unwrap();
        std::fs::File::open(&f_path).unwrap().sync_data().unwrap();
    }

    fn get_tmp_file_path(&self) -> PathBuf {
        self.root_storage_dir.join(Storage::tmp_dir()).join("tmp")
    }

    fn validate_key(key: &str) -> Result<(), String> {
        if key.len() <= 255 {
            return Ok(());
        }
        Err("Key is too long".to_string())
    }

    fn validate_value(value: &[u8]) -> Result<(), String> {
        if value.len() <= 65535 {
            return Ok(());
        }
        return Err("Value is too big".to_string());
    }

    fn get_file_path(&self, path: &str) -> PathBuf {
        return self.root_storage_dir.join(path);
    }

    fn get_key_line(key: &str) -> String {
        format!("{}{}\n", key.len() as u8 as char, key)
    }

    async fn save_in_key_storage(&self, key: &str) {
        let mut f = OpenOptions::new()
            .append(true)
            .open(self.get_file_path(KEYS_FILE_NAME))
            .await
            .unwrap();

        let data = Storage::get_key_line(key);
        f.write_all(data.as_bytes()).await.unwrap();
        f.sync_data().await.unwrap();
    }

    async fn save_to_file_atomically(&self, name: &str, value: &[u8]) {
        let tmp_path = self.get_tmp_file_path();

        let mut file = tokio::fs::File::create(&tmp_path).await.unwrap();
        file.write_all(value).await.unwrap();
        file.sync_data().await.unwrap();

        let save_path = self.get_file_path(&name);

        tokio::fs::rename(&tmp_path, &save_path).await.unwrap();
        tokio::fs::File::open(save_path)
            .await
            .unwrap()
            .sync_data()
            .await
            .unwrap();
    }
}

#[async_trait::async_trait]
impl StableStorage for Storage {
    async fn put(&mut self, key: &str, value: &[u8]) -> Result<(), String> {
        if let Err(msg) = Storage::validate_key(key) {
            return Err(msg);
        }
        if let Err(msg) = Storage::validate_value(value) {
            return Err(msg);
        }
        match self.name_map.get(key) {
            None => {
                let order = self.name_map.len().to_string();

                self.save_to_file_atomically(&order, value).await;

                self.save_in_key_storage(key).await;
                self.name_map.insert(key.to_string(), order);

                Ok(())
            }
            Some(name) => {
                self.save_to_file_atomically(name, value).await;

                Ok(())
            }
        }
    }

    async fn get(&self, key: &str) -> Option<Vec<u8>> {
        println!("{}", self.name_map.len());

        if let Err(_) = Storage::validate_key(key) {
            return None;
        }

        match self.name_map.get(key) {
            None => None,
            Some(name) => {
                let value = tokio::fs::read_to_string(self.get_file_path(&name))
                    .await
                    .unwrap()
                    .as_bytes()
                    .to_vec();
                Some(value)
            }
        }
    }
}

/// Creates a new instance of stable storage.
pub async fn build_stable_storage(root_storage_dir: PathBuf) -> Box<dyn StableStorage> {
    tokio::task::spawn_blocking(|| {
        let mut storage = Storage::new(root_storage_dir);
        storage.restore();
        Box::new(storage)
    })
    .await
    .unwrap()
}
