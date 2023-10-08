#include <iomanip>
#include <iostream>
#include <span>
#include <sstream>
#include <unordered_map>
#include <unordered_set>
#include <vector>

#include <capstone/capstone.h>
#include <keystone/keystone.h>

#include "elfio/elfio.hpp"

namespace conv {

const std::vector<std::string> ARM_PROLOG = {"stp x29, x30, [sp, #-16]!", "mov x29, sp"};
const std::vector<std::string> ARM_EPILOG = {"mov x0, x9", "add sp, x29, #16", "ldp x29, x30, [sp, #-16]", "ret"};

const std::vector<std::string> X86_64_PROLOG = {"endbr64", "push rbp", "mov rbp, rsp"};
const std::vector<std::string> X86_64_EPILOG = {"leave", "ret"};

std::string trim(const std::string& s) {
  auto start = s.find_first_not_of(' ');
  auto end = s.find_last_not_of(' ');
  return s.substr(start, end - start + 1);
}

std::string str_instr(const cs_insn* insn, bool address = false) {
  std::string str_in;
  if (address) {
    str_in = std::to_string(insn->address) + ": ";
  }
  str_in += std::string(insn->mnemonic) + " " + std::string(insn->op_str);
  return trim(str_in);
}

void check_prolog(const cs_insn* insn, size_t count) {
  if (count < X86_64_PROLOG.size()) {
    throw std::runtime_error("Function do not have a valid prolog");
  }
  for (size_t i = 0; i < X86_64_PROLOG.size(); i++) {
    std::string s = str_instr(&insn[i]);
    if (s != X86_64_PROLOG[i]) {
      throw std::runtime_error("Function do not have a valid prolog");
    }
  }
}

void check_epilog(const cs_insn* insn, size_t count) {
  if (count < X86_64_PROLOG.size() + X86_64_EPILOG.size()) {
    throw std::runtime_error("Function do not have a valid epilog");
  }
  for (size_t i = 0; i < X86_64_EPILOG.size(); i++) {
    if (str_instr(&insn[count - X86_64_EPILOG.size() + i]) != X86_64_EPILOG[i]) {
      throw std::runtime_error("Function do not have a valid epilog");
    }
  }
}

const std::string ARM_TMP1_64 = "x12";
const std::string ARM_TMP1_32 = "w12";
const std::string ARM_TMP2_64 = "x13";
const std::string ARM_TMP_32 = "w13";

const std::unordered_map<x86_insn, std::string> JUMP_MAP = {
    {X86_INS_JMP, "b"},     //
    {X86_INS_JA, "b.hi"},   // ja, jnbe
    {X86_INS_JAE, "b.hs"},  // jae, jnb
    {X86_INS_JB, "b.lo"},   // jb, jnae
    {X86_INS_JBE, "b.ls"},  // jbe, jna
    {X86_INS_JE, "b.eq"},   // je, jz
    {X86_INS_JG, "b.gt"},   //jg, jnle
    {X86_INS_JGE, "b.ge"},  // jge, jnl
    {X86_INS_JL, "b.lt"},   //
    {X86_INS_JLE, "b.le"},  // jle, jng!
    {X86_INS_JNE, "b.ne"},  // jne, jnz
    {X86_INS_JNO, "b.vc"},  //
    {X86_INS_JO, "b.vs"},
};

const std::unordered_map<x86_reg, std::string> REGISTERS_MAP_64 = {
    //Caller-saved:
    {X86_REG_RDI, "x0"},  // 1. argument
    {X86_REG_RSI, "x1"},  // 2. argument
    {X86_REG_RDX, "x2"},  // 3. argument
    {X86_REG_RCX, "x3"},  // 4. argument
    {X86_REG_R8, "x4"},   // 5. argument
    {X86_REG_R9, "x5"},   // 6. argument
    {X86_REG_RAX, "x9"},
    {X86_REG_R10, "x10"},
    {X86_REG_R11, "x11"},
    // Callee-saved:
    {X86_REG_RBP, "x29"},  // fp
    {X86_REG_RBX, "x19"},
    {X86_REG_R12, "x20"},
    {X86_REG_R13, "x21"},
    {X86_REG_R14, "x22"},
    {X86_REG_R15, "x23"},
    {X86_REG_RSP, "sp"},
};

const std::unordered_map<x86_reg, std::string> REGISTERS_MAP_32 = {
    // Caller-saved:
    {X86_REG_EDI, "w0"},  // 1. argument
    {X86_REG_ESI, "w1"},  // 2. argument
    {X86_REG_EDX, "w2"},  // 3. argument
    {X86_REG_ECX, "w3"},  // 4. argument
    {X86_REG_R8D, "w4"},  // 5. argument
    {X86_REG_R9D, "w5"},  // 6. argument
    {X86_REG_EAX, "w9"},
    {X86_REG_R10D, "w10"},
    {X86_REG_R11D, "w11"},
    // Callee-saved:
    {X86_REG_EBP, "w29"},  // fp
    {X86_REG_EBX, "w19"},
    {X86_REG_R12D, "w20"},
    {X86_REG_R13D, "w21"},
    {X86_REG_R14D, "w22"},
    {X86_REG_R15D, "w23"},
};

bool is_64bit(cs_x86_op op) {
  return op.size == 8;
}

std::string convert_register(x86_reg reg) {
  auto it = REGISTERS_MAP_32.find(reg);
  if (it == REGISTERS_MAP_32.end()) {
    return REGISTERS_MAP_64.at(reg);
  }
  return it->second;
}

std::string convert_operand(cs_x86_op op) {
  switch (op.type) {
    case X86_OP_REG:
      return convert_register(op.reg);
    case X86_OP_IMM:
      return std::to_string(op.imm);
    default:
      throw std::runtime_error("Unexpected operand type" + std::to_string(op.type));
  }
}

std::string to_label(uint64_t from_addr) {
  return "l" + std::to_string(from_addr);
}

std::string to_64bit_arm_reg(const std::string& reg) {
  if (reg[0] == 'w') {
    return "x" + reg.substr(1);
  }
  return reg;
}

std::unordered_map<uint64_t, uint64_t> get_labels(const cs_insn* insn, size_t count) {
  std::unordered_map<uint64_t, uint64_t> labels;
  for (size_t i = 0; i < count; i++) {
    if (JUMP_MAP.contains(static_cast<x86_insn>(insn[i].id))) {
      labels.insert({insn[i].detail->x86.operands[0].imm, insn[i].address});
    }
  }
  return labels;
}

struct Symbol {
  ELFIO::Elf_Xword index;
  ELFIO::Elf_Word name;
  ELFIO::Elf64_Addr value;
  ELFIO::Elf_Xword size;
  unsigned char bind;
  unsigned char type;
  unsigned char other;
  [[nodiscard]] bool isZero() const {
    return index == 0 && name == 0 && value == 0 && size == 0 && bind == 0 && type == 0 && other == 0;
  }
};

struct Relocation {
  ELFIO::Elf_Xword index;
  ELFIO::Elf64_Addr offset;
  ELFIO::Elf_Word symbol;
  unsigned type;
  ELFIO::Elf_Sxword addend;
};

class ASMConverter {
  char instr_end() {
    instr_count++;
    return '\n';
  }

  std::stringstream ss;
  size_t instr_count{};
  ELFIO::Elf64_Addr start;

  std::vector<Relocation>& new_relocs;
  std::vector<Relocation>::const_iterator& curr_reloc;
  const std::vector<Relocation>::const_iterator& relocs_end;

  void mem_load(const std::string& dst, const cs_insn* instr, cs_x86_op op) {
    x86_reg reg = op.mem.base;
    switch (reg) {
      case X86_REG_RIP:
        convert_relocation(ELFIO::R_AARCH64_LD_PREL_LO19, instr);
        ss << "ldr " << dst << ", #0" << instr_end();
        return;
      default:
        ss << "mov " << ARM_TMP1_64 << ", " << op.mem.disp << instr_end();
        ss << "ldr " << dst << ", [" << convert_register(reg) << ", " << ARM_TMP1_64 << "]" << instr_end();
        return;
    }
  }

 public:
  explicit ASMConverter(std::vector<Relocation>& _new_relocs, std::vector<Relocation>::const_iterator& _curr_reloc,
                        const std::vector<Relocation>::const_iterator& _relocs_end, ELFIO::Elf64_Addr _start)
      : new_relocs(_new_relocs), curr_reloc(_curr_reloc), relocs_end(_relocs_end), start(_start) {}

  void add_instr(const std::string& instr) { ss << instr << instr_end(); }

  void convert_relocation(unsigned rel_type, const cs_insn* instr) {
    // FIXME currently it is only a happy path not checking if it is in fact a valid relocation.
    Relocation relocation{};
    switch (rel_type) {
      case ELFIO::R_AARCH64_LD_PREL_LO19:
        switch (curr_reloc->type) {
          case ELFIO::R_X86_64_PC32:
          case ELFIO::R_X86_64_PLT32: {
            relocation = *curr_reloc;
            relocation.addend += instr->size;  // RIP points to the next instruction.
            // Previously offset included position of the operand in the instruction.
            relocation.addend -= static_cast<ELFIO::Elf_Sxword>(relocation.offset - instr->address);
          } break;
          default:
            throw std::runtime_error("unexpected relocation type: " + std::to_string(curr_reloc->type));
        }
        break;
      case ELFIO::R_AARCH64_ADR_PREL_LO21:
        switch (curr_reloc->type) {
          // mov [rip + X], reg/imm
          case ELFIO::R_X86_64_PC32:
          case ELFIO::R_X86_64_PLT32:
            relocation = *curr_reloc;
            relocation.addend += instr->size;  // RIP points to the next instruction.
            // Previously offset included position of the operand in the instruction.
            relocation.addend -= static_cast<ELFIO::Elf_Sxword>(relocation.offset - instr->address);
            break;
            // mov mem, imm(with reloc)
            // mov reg, imm(with reloc)
          case ELFIO::R_X86_64_32:
          case ELFIO::R_X86_64_32S:
            relocation = *curr_reloc;
            break;
          default:
            throw std::runtime_error("unexpected relocation type: " + std::to_string(curr_reloc->type));
        }
        break;
      case ELFIO::R_AARCH64_CALL26:
        switch (curr_reloc->type) {
          case ELFIO::R_X86_64_PC32:
          case ELFIO::R_X86_64_PLT32:
            relocation = *curr_reloc;
            relocation.addend += 4;  // Branch in arm do not need to adjust offset as x86 has to in call.
            break;
          default:
            throw std::runtime_error("unexpected relocation type: " + std::to_string(curr_reloc->type));
        }
        break;
      default:
        throw std::runtime_error("unexpected relocation type: " + std::to_string(curr_reloc->type));
    }
    relocation.type = rel_type;
    relocation.offset = instr_count * 4 + start;  // Each ARM instruction has 4 bytes.

    new_relocs.push_back(relocation);

    curr_reloc++;
  }

  bool check_mov_relocation(const cs_insn* insn) {
    return (curr_reloc != relocs_end) && (insn->detail->x86.operands[1].type == X86_OP_IMM) &&
           ((curr_reloc->type == ELFIO::R_X86_64_32) || (curr_reloc->type == ELFIO::R_X86_64_32S)) &&
           ((insn->address <= curr_reloc->offset) && (curr_reloc->offset < insn->address + insn->size));
  }

  void convert_instr(const cs_insn* instr) {
    cs_x86 instr_x86 = instr->detail->x86;
    auto type = static_cast<x86_insn>(instr->id);

    auto jmp_it = JUMP_MAP.find(type);
    if (jmp_it != JUMP_MAP.end()) {
      ss << jmp_it->second << ' ' << to_label(instr->address) << instr_end();
      return;
    }

    switch (type) {
      case X86_INS_SUB:
      case X86_INS_ADD: {
        std::string op1 = convert_operand(instr_x86.operands[0]);
        std::string op2 = convert_operand(instr_x86.operands[1]);

        std::string mnemonic = type == X86_INS_SUB ? "sub" : "add";

        ss << mnemonic << ' ' << op1 << ", " << op1 << ", " << op2 << instr_end();
        break;
      }
      case X86_INS_CMP: {
        cs_x86_op op1 = instr_x86.operands[0];
        cs_x86_op op2 = instr_x86.operands[1];

        std::string n_op1, n_op2;

        switch (op1.type) {
          case X86_OP_MEM: {
            n_op1 = is_64bit(op1) ? ARM_TMP1_64 : ARM_TMP1_32;
            mem_load(n_op1, instr, op1);
            n_op2 = convert_operand(op2);
            break;
          }
          case X86_OP_REG: {
            n_op1 = convert_operand(op1);
            switch (op2.type) {
              case X86_OP_REG:
              case X86_OP_IMM:
                n_op2 = convert_operand(op2);
                break;
              case X86_OP_MEM:
                n_op2 = is_64bit(op2) ? ARM_TMP1_64 : ARM_TMP1_32;
                mem_load(n_op2, instr, op2);
                break;
              default:
                throw std::runtime_error("Unsupported type of second operand in cmp " + std::to_string(op1.type));
            }
            break;
          }
          default:
            throw std::runtime_error("Unsupported type of first operand in cmp " + std::to_string(op1.type));
        }

        ss << "cmp " << n_op1 << ", " << n_op2 << instr_end();
        break;
      }
      case X86_INS_CALL: {
        convert_relocation(ELFIO::R_AARCH64_CALL26, instr);
        ss << "bl #0" << instr_end();
        ss << "mov x9, x0" << instr_end();
        break;
      }
      case X86_INS_MOV: {
        cs_x86_op op1 = instr_x86.operands[0];
        cs_x86_op op2 = instr_x86.operands[1];

        switch (op1.type) {
          case X86_OP_MEM: {
            if (op1.reg == X86_REG_RIP) {
              std::string tmp_reg = is_64bit(op1) ? ARM_TMP2_64 : ARM_TMP_32;
              ss << "adr " << ARM_TMP1_64 << ", #0" << instr_end();
              ss << "mov " << tmp_reg << convert_operand(op2) << instr_end();
              ss << "str " << tmp_reg << ", [" << ARM_TMP1_64 << "]" << instr_end();
            } else {
              std::string base = convert_register(op1.mem.base);
              if (check_mov_relocation(instr)) {
                convert_relocation(ELFIO::R_AARCH64_ADR_PREL_LO21, instr);
                ss << "adr " << ARM_TMP1_64 << ", #0" << instr_end();
                ss << "mov " << ARM_TMP2_64 << ", #" << op1.mem.disp << instr_end();
                ss << "str " << ARM_TMP1_32 << ", [" << base << ", " << ARM_TMP2_64 << "]" << instr_end();
              } else {
                std::string tmp1 = is_64bit(op2) ? ARM_TMP1_64 : ARM_TMP1_32;
                ss << "mov " << tmp1 << ", " << convert_operand(op2) << instr_end();
                ss << "mov " << ARM_TMP2_64 << ", #" << op1.mem.disp << instr_end();
                ss << "str " << tmp1 << ", [" << base << ", " << ARM_TMP2_64 << "]" << instr_end();
              }
            }
            break;
          }
          case X86_OP_REG: {
            switch (op2.type) {
              case X86_OP_IMM:
              case X86_OP_REG:
                if (check_mov_relocation(instr)) {
                  convert_relocation(ELFIO::R_AARCH64_ADR_PREL_LO21, instr);
                  ss << "adr " << to_64bit_arm_reg(convert_register(op1.reg)) << ", #0" << instr_end();
                  break;
                }
                ss << "mov " << convert_register(op1.reg) << ", " << convert_operand(op2) << instr_end();
                break;
              case X86_OP_MEM:
                mem_load(convert_register(op1.reg), instr, op2);
                break;
              default:
                throw std::runtime_error("Unsupported type of second operand in cmp " + std::to_string(op1.type));
            }
            break;
          }
          default:
            throw std::runtime_error("Unsupported type of first operand in mov " + std::to_string(op1.type));
        }
        break;
      }
      default:
        throw std::runtime_error("Unexpected instruction " + std::string(instr->mnemonic) + " " +
                                 std::string(instr->op_str));
    }
  }

  void add_label(uint64_t from_addr) { ss << to_label(from_addr) << ":\n"; }

  std::string get() const { return ss.str(); }
};

class Builder {
  ELFIO::elfio elf;
  std::unordered_map<ELFIO::Elf_Half, ELFIO::Elf_Half> shndx_mapping;
  std::vector<std::pair<Symbol, ELFIO::Elf_Half>> symbols;

  static ELFIO::section* get_symbol_section(const ELFIO::elfio& org_elf) {
    for (auto& section : org_elf.sections) {
      if (section->get_type() == ELFIO::SHT_SYMTAB) {
        return section.get();
      }
    }
    throw std::runtime_error("symbol section not found");
  }

  void adjust_section_addresses() {
    ELFIO::Elf64_Addr addr = 0;
    for (auto& section : elf.sections) {
      section->set_address(addr);
      addr += section->get_size();
    }
  }

  void adjust_section_refs() {
    for (auto& section : elf.sections) {
      ELFIO::Elf_Word link = section->get_link();
      if (link != 0) {
        section->set_link(shndx_mapping[link]);
      }
      if (section->get_type() == ELFIO::SHT_RELA) {
        ELFIO::Elf_Word info = section->get_info();
        section->set_info(shndx_mapping[info]);
      }
    }
    elf.set_section_name_str_index(shndx_mapping[elf.get_section_name_str_index()]);
  }

  void store_symbols() {
    if (symbols.empty()) {
      return;
    }
    ELFIO::section* section = get_symbol_section(elf);
    ELFIO::symbol_section_accessor sym_tab(elf, section);

    std::sort(symbols.begin(), symbols.end(),
              [](const auto& a, const auto& b) -> bool { return a.first.index < b.first.index; });

    ELFIO::Elf_Xword expected_idx = 0;
    ELFIO::Elf_Xword expected_bind = ELFIO::STB_LOCAL;

    for (auto [symbol, shndx] : symbols) {
      if (shndx == 0 && symbol.isZero()) {
        expected_idx = 1;
        continue;  // ELFIO generated null symbol by itself;
      }
      for (; expected_idx < symbol.index; expected_idx++) {
        // Add dummy symbols to replace symbols from removed sections.
        sym_tab.add_symbol(0, 0, 0, 0, 0, 0, 0);
      }
      expected_idx = symbol.index + 1;
      expected_bind = symbol.bind;

      auto it = shndx_mapping.find(shndx);
      if (it != shndx_mapping.end()) {
        shndx = it->second;
      }
      sym_tab.add_symbol(symbol.name, symbol.value, symbol.size, symbol.bind, symbol.type, symbol.other, shndx);
    }
  }

  static ELFIO::elfio prepare_elf(unsigned char file_class, unsigned char encoding, ELFIO::Elf_Half machine,
                                  const ELFIO::elfio& org_elf) {
    ELFIO::elfio elf;
    elf.create_no_sec(file_class, encoding);
    elf.set_section_name_str_index(org_elf.get_section_name_str_index());
    elf.set_machine(machine);
    elf.set_type(org_elf.get_type());
    elf.set_os_abi(org_elf.get_os_abi());
    elf.set_flags(org_elf.get_flags());

    return elf;
  }

 public:
  Builder(unsigned char file_class, unsigned char encoding, ELFIO::Elf_Half machine, const ELFIO::elfio& org_elf)
      : elf(prepare_elf(file_class, encoding, machine, org_elf)){};

  void save(const std::string& path) {
    store_symbols();
    adjust_section_addresses();
    adjust_section_refs();
    if (!elf.save(path)) {
      throw std::runtime_error("Can't save converter file at: " + path);
    }
  };

  ELFIO::section* copy_section(const ELFIO::section* section, bool copy_data = false) {
    auto new_sec = elf.sections.add(section->get_name_string_offset());

    shndx_mapping.insert({section->get_index(), new_sec->get_index()});

    new_sec->set_addr_align(section->get_addr_align());
    if (copy_data) {
      const char* data = section->get_data();
      if (data != nullptr) {
        new_sec->set_data(data);
      }
    }
    new_sec->set_entry_size(section->get_entry_size());
    new_sec->set_flags(section->get_flags());
    new_sec->set_info(section->get_info());
    new_sec->set_link(section->get_link());
    new_sec->set_type(section->get_type());

    return new_sec;
  }

  void add_symbol(Symbol symbol, ELFIO::Elf_Half shndx) { symbols.emplace_back(symbol, shndx); }
};

class Converter {
 public:
  explicit Converter(ELFIO::elfio&& _elf)
      : elf(std::move(_elf)),
        sym_tab(get_sym_tab()),
        section_symbols(get_section_symbols()),
        section_relocations(get_section_relocations()),
        builder(ELFIO::ELFCLASS64, ELFIO::ELFDATA2LSB, ELFIO::EM_AARCH64, elf) {
    if (cs_open(CS_ARCH_X86, CS_MODE_64, &handle) != CS_ERR_OK) {
      throw std::runtime_error("Error starting capstone");
    }
    cs_option(handle, CS_OPT_DETAIL, CS_OPT_ON);
    if (ks_open(KS_ARCH_ARM64, KS_MODE_LITTLE_ENDIAN, &ks) != KS_ERR_OK) {
      throw std::runtime_error("Failed to open keystone");
    }
  }

  ~Converter() {
    cs_close(&handle);
    ks_close(ks);
  }

  void convert(const std::string& path) {
    convert_code();

    builder.save(path);
  }

 private:
  csh handle{};
  ks_engine* ks{};

  const ELFIO::elfio elf;
  const ELFIO::const_symbol_section_accessor sym_tab;
  const std::unordered_map<ELFIO::Elf_Half, std::vector<Symbol>> section_symbols;
  const std::unordered_map<ELFIO::Elf_Half, std::pair<const ELFIO::section*, std::vector<Relocation>>>
      section_relocations;
  Builder builder;

  std::unordered_map<ELFIO::Elf_Half, std::vector<Symbol>> get_section_symbols() {
    std::unordered_map<ELFIO::Elf_Half, std::vector<Symbol>> syms;

    for (ELFIO::Elf_Xword i = 0; i < sym_tab.get_symbols_num(); i++) {
      Symbol symbol{.index = i};
      ELFIO::Elf_Half section_index;

      sym_tab.get_symbol2(i, symbol.name, symbol.value, symbol.size, symbol.bind, symbol.type, section_index,
                          symbol.other);  // Gotta love this API.
      auto it = syms.find(section_index);
      if (it == syms.end()) {
        syms.insert({section_index, {symbol}});
      } else {
        it->second.push_back(symbol);
      }
    }

    // Sort symbols by addr, but function at the end in case of multiple symbols with the same addr..
    for (auto& kv : syms) {
      std::sort(kv.second.begin(), kv.second.end(), [](const Symbol& a, const Symbol& b) -> bool {
        if (a.value == b.value) {
          return a.type != ELFIO::STT_FUNC;
        }
        return a.value < b.value;
      });
    }

    return syms;
  }

  std::unordered_map<ELFIO::Elf_Half, std::pair<const ELFIO::section*, std::vector<Relocation>>>
  get_section_relocations() {
    std::unordered_map<ELFIO::Elf_Half, std::pair<const ELFIO::section*, std::vector<Relocation>>> relocs;

    for (auto& section : elf.sections) {
      if (section->get_type() != ELFIO::SHT_RELA) {
        continue;
      }
      ELFIO::Elf_Half section_index = section->get_info();

      const ELFIO::relocation_section_accessor rel_tab{elf, section.get()};
      for (ELFIO::Elf_Xword i = 0; i < rel_tab.get_entries_num(); i++) {
        Relocation reloc{.index = i};
        rel_tab.get_entry(i, reloc.offset, reloc.symbol, reloc.type, reloc.addend);

        auto it = relocs.find(section_index);
        if (it == relocs.end()) {
          relocs.insert({section_index, {section.get(), {reloc}}});
        } else {
          it->second.second.push_back(reloc);
        }
      }
    }

    // Sort relocs by offs.
    for (auto& kv : relocs) {
      std::sort(kv.second.second.begin(), kv.second.second.end(),
                [](const Relocation& a, const Relocation& b) -> bool { return a.offset < b.offset; });
    }

    return relocs;
  }

  ELFIO::const_symbol_section_accessor get_sym_tab() {
    ELFIO::section* sec{};
    for (const auto& section : elf.sections) {
      if (section->get_type() == ELFIO::SHT_SYMTAB) {
        if (sec != nullptr) {
          throw std::runtime_error("Found more than one symbol table sections.");
        }
        sec = section.get();
      }
    }
    if (sec == nullptr) {
      throw std::runtime_error("Didn't find a symbol table section.");
    }

    return ELFIO::const_symbol_section_accessor(elf, sec);
  }

  std::string convert_fn(const std::string_view& code, std::vector<Relocation>& new_relocs,
                         std::vector<Relocation>::const_iterator& curr_reloc,
                         const std::vector<Relocation>::const_iterator& end_relocs, ELFIO::Elf64_Addr start) {
    cs_insn* insn;
    size_t count;

    count = cs_disasm(handle, reinterpret_cast<const uint8_t*>(code.data()), code.size(), 0, 0, &insn);
    if (count == 0) {
      throw std::runtime_error("Failed to disassemble given code!");
    }

    std::span<cs_insn> instrs{insn, count};  // For debug

    check_prolog(insn, count);
    check_epilog(insn, count);

    ASMConverter asm_converter{new_relocs, curr_reloc, end_relocs, start};

    std::unordered_map<uint64_t, uint64_t> labels = get_labels(insn, count);

    for (const auto& ins : ARM_PROLOG) {
      asm_converter.add_instr(ins);
    }

    for (size_t i = X86_64_PROLOG.size(); i < count - X86_64_EPILOG.size(); i++) {
      auto label_it = labels.find(insn[i].address);
      if (label_it != labels.end()) {
        asm_converter.add_label(label_it->second);
      }

      asm_converter.convert_instr(&insn[i]);
    }

    // Add label just before leave.
    auto label_it = labels.find(insn[count - X86_64_EPILOG.size()].address);
    if (label_it != labels.end()) {
      asm_converter.add_label(label_it->second);
    }

    for (const auto& ins : ARM_EPILOG) {
      asm_converter.add_instr(ins);
    }

    cs_free(insn, count);

    std::string arm_code = asm_converter.get();

    std::cerr << arm_code << std::endl;

    std::string new_code = assemble(arm_code);

    return new_code;
  }

  std::vector<Relocation> convert_section(ELFIO::section* section, const std::string_view& data,
                                          const std::vector<Symbol>& symbols,
                                          const std::vector<Relocation>& relocations) {
    ELFIO::Elf64_Addr prev_fn_end = 0;
    std::string new_data;
    ssize_t diff = 0;

    auto curr_reloc = relocations.begin();
    auto end_relocs = relocations.end();

    std::vector<Relocation> new_relocs;

    for (auto symbol : symbols) {
      // Out of function relocations.
      while (curr_reloc != end_relocs && curr_reloc->offset < symbol.value) {
        if (curr_reloc->type != ELFIO::R_X86_64_64) {
          throw std::runtime_error("Unexpected out of function relocations: " + std::to_string(curr_reloc->type));
        }

        Relocation reloc = *curr_reloc;
        reloc.type = ELFIO::R_AARCH64_ABS64;
        reloc.offset += diff;

        new_relocs.push_back(reloc);

        curr_reloc++;
      }

      switch (symbol.type) {
        case ELFIO::STT_FUNC: {
          std::string new_fn = convert_fn(data.substr(symbol.value, symbol.size), new_relocs, curr_reloc, end_relocs,
                                          symbol.value + diff);

          new_data += data.substr(prev_fn_end, symbol.value - prev_fn_end);
          new_data += new_fn;

          prev_fn_end = symbol.value + symbol.size;

          // Adjust symbol and diff.
          symbol.value += diff;
          diff += ssize_t(new_fn.length()) - ssize_t(symbol.size);
          symbol.size = new_fn.length();

          break;
        }

        case ELFIO::STT_OBJECT:
        case ELFIO::STT_NOTYPE: {
          symbol.value += diff;
          break;
        }
      }
      builder.add_symbol(symbol, section->get_index());
    }

    // Out of function relocations.
    while (curr_reloc != end_relocs) {
      if (curr_reloc->type != ELFIO::R_X86_64_64) {
        throw std::runtime_error("Unexpected out of function relocations: " + std::to_string(curr_reloc->type));
      }

      Relocation reloc = *curr_reloc;
      reloc.type = ELFIO::R_AARCH64_ABS64;
      reloc.offset += diff;

      new_relocs.push_back(reloc);

      curr_reloc++;
    }

    new_data += data.substr(prev_fn_end, data.size() - prev_fn_end);

    section->set_data(new_data);
    return new_relocs;
  }

  std::string assemble(const std::string& code) {
    size_t count;
    unsigned char* encode;
    size_t size;

    if (ks_asm(ks, code.c_str(), 0, &encode, &size, &count) != KS_ERR_OK) {
      throw std::runtime_error("ks_asm() failed & count = " + std::to_string(count) +
                               ", error = " + std::string(ks_strerror(ks_errno(ks))));
    }

    std::string compiled(reinterpret_cast<char*>(encode), size);
    ks_free(encode);
    return compiled;
  }

  void convert_section_relocations(ELFIO::section* relocs_section, const std::vector<Relocation>& new_relocs) {
    ELFIO::relocation_section_accessor relocs{elf, relocs_section};

    for (Relocation const& reloc : new_relocs) {
      relocs.add_entry(reloc.offset, reloc.symbol, reloc.type, reloc.addend);
    }
  }

  void convert_code() {
    const std::pair<ELFIO::section*, std::vector<Relocation>> dummy_relocs{};
    const std::vector<Symbol> dummy_symbols{};

    for (const auto& section : elf.sections) {
      if (ignore_section(section.get())) {
        continue;
      }
      if (section->get_type() == ELFIO::SHT_SYMTAB) {
        // Copy symtable without data.
        builder.copy_section(section.get(), true);
        continue;
      }

      auto sec_copy = builder.copy_section(section.get());

      auto symbols_it = section_symbols.find(section->get_index());
      auto& symbols = symbols_it == section_symbols.end() ? dummy_symbols : symbols_it->second;
      const auto& relocs_it = section_relocations.find(section->get_index());
      auto& [relocs_sec, relocs] = relocs_it == section_relocations.end() ? dummy_relocs : relocs_it->second;

      auto new_relocs = convert_section(sec_copy, {section->get_data(), section->get_size()}, symbols, relocs);

      if (relocs_sec != nullptr) {
        ELFIO::section* relocs_sec_copy = builder.copy_section(relocs_sec);
        convert_section_relocations(relocs_sec_copy, new_relocs);
      }
    }

    // Add remaining symbols.
    for (const auto& kv : section_symbols) {
      if (kv.first < elf.sections.size()) {
        continue;
      }
      for (const auto& sym : kv.second) {
        builder.add_symbol(sym, kv.first);
      }
    }
  }

  static bool ignore_section(const ELFIO::section* section) {
    std::string name = section->get_name();
    return (name == ".note.gnu.property" || name.ends_with(".eh_frame")) ||  // by name
           (section->get_type() == ELFIO::SHT_RELA);
  }
};

void convert(const std::string& in_path, const std::string& out_path) {
  ELFIO::elfio elf;

  if (!elf.load(in_path)) {
    throw std::runtime_error("Can't find or process ELF file " + in_path);
  }

  Converter converter{std::move(elf)};
  converter.convert(out_path);
}
}  // namespace conv

int main(int argc, char** argv) {
  if (argc != 3) {
    std::cerr << "Usage: " << argv[0] << "<path/to/converted/elf> <path/to/output/elf>" << std::endl;
    return 1;
  }

  conv::convert(argv[1], argv[2]);
  return 0;
}