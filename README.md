# clean_pdb_waters
`clean_pdb_waters` is a simple command-line tool written in Go that removes all water molecules (HOH) with zero occupancy from a PDB file.

## 🧠 What It Does

This tool scans a PDB file line by line and removes any atom records corresponding to water molecules (`HOH`) that have an occupancy value of `0.00`. The cleaned structure is saved to a new file.

## 📦 Requirements

- Go 1.18 or newer

## 🔧 Installation

Clone the repository and build the program:

git clone https://github.com/kolenpe1/clean_pdb_waters.git
cd clean_pdb_waters
go build -o clean_pdb_waters

## 🤝 Acknowledgments
Special thanks to Microsoft Copilot for helping with code generation and cleanup logic.
