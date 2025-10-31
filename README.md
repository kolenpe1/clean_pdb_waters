# clean_waters
`clean_waters` is a simple command-line tool written in Go that removes all water molecules (HOH) with zero occupancy from a PDB, CIF or mmCIF file.

## 🧠 What It Does

This tool scans structure file line by line and removes any atom records corresponding to water molecules (`HOH`) that have an occupancy value of `0.00`. The cleaned structure is saved to a new file.

In final stages of structure refinement, questionable waters are repeatedly added and removed. Waters with zero occupancy can be used as flags for maxima in electron density that will not be interpreted.

## 📦 Requirements

- Binary file is standalone
- Go 1.18 or newer for compilation from the source

## 🔧 Installation

Clone the repository and build the program:

```bash
git clone https://github.com/kolenpe1/clean_waters.git
cd clean_waters
go build -o clean_waters
```

Copy file 'clean_waters' to your 'bin' directory and run.


## 🧪 Example

To clean a PDB file named `example.pdb`, run:
```bash
clean_waters example.pdb
```

This will produce `cleaned_example.pdb` without the waters with zero occupancy.


## 🤝 Acknowledgments
Special thanks to Microsoft Copilot for helping with code generation and cleanup logic.
