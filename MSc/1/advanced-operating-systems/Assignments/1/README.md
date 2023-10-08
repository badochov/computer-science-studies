# ELF converter from x64 to arm64 

[Detailed description](https://students.mimuw.edu.pl/ZSO/PUBLIC-SO/2022-2023/z1_elf/index.html#zadanie)

## Solution

The solution uses 3 external libraries:
-  capstone
  - for disassembly
- keystone
  - for assembly
- elfio
  - included in the package in ELFIO/ folder
  - to read x64 elf and write arm64
  - it had to be slightly modified as some functions wanted to do too much
      - diff from the modifications can be seen in [patch](ELFIO/0001-Functions-without-unneccessary-copies.patch).

Flow of the program:
- parsing arguments
- reading elf
- new elf creation - setting appropriate header
- getting information about each section's symbols and relocations
- converting each section that is not `SHT_RELA`
  - going through each symbol and out of function relocation in order of offsets ascending one by one updating offsets
  - if the symbol is a function:
    - function disassembly using capstone
    - obtaining all jump addresses to know when to insert labels
    - converting its instructions one by one according to conversion instructions
    - if instruction has a relocation converting the relocation to appropriate type
    - function assembly using keystone
  - saving updated contents of the section
  - if needed creation of `SHT_RELA` containing section's relocations
- adjusting offsets of sections according to new sizes
- adjusting section indexes according to new indexes
- saving newly generated elf file

## Compilation
For Debian on provided qemu:
### Dependencies
- cmake
- make
- capstone
```bash
sudo apt update
sudo apt install cmake make libcapstone-dev
```
- keystone
```bash
wget https://github.com/keystone-engine/keystone/archive/refs/tags/0.9.2.tar.gz
tar -xf 0.9.2.tar.gz
cd keystone-0.9.2
mkdir build 
cd build
../make-share.sh
sudo make install
sudo ldconfig
```
### Compilation
```
cd <path/to/converter/dir>
mkdir build 
cd build
cmake ..
make
```


## Usage
```
./converter <path/to/x64/rela/elf> <path/to/out/arm64/rela/elf>
```
