# SnapArchive

An application that recursively reads files from a source directory and copies them to a destination directory organized by year and month based on the file's modification time.
Mainly to organize my photos, but can work on any type of files.

## Features

- Recursively processes all files in the source directory and its subdirectories
- Organizes files by year and month (YYYY/YYYY-MM-DD format)
- Multiple file conflict resolution strategies: skip, overwrite, or rename
- Preserves file permissions
- Provides clear error messages and validation
- Shows progress during file copying

## Usage

```bash
./snaparchive [options] <source_folder> <dest_folder>
```

### Arguments

- `source_folder`: The path to the directory containing files to organize
- `dest_folder`: The path where organized files will be copied

### Options

- `-overwrite <mode>`: Behavior when file exists (default: "skip")
  - `skip`: Skip files that already exist in destination (default)
  - `overwrite`: Replace existing files with source files
  - `rename`: Create new file with incremented number suffix

### Examples

#### Basic usage (default skip behavior)
```bash
./snaparchive ./photos ./organized_photos
```

#### Overwrite existing files
```bash
./snaparchive -overwrite overwrite ./photos ./organized_photos
```

#### Rename conflicting files
```bash
./snaparchive -overwrite rename ./photos ./organized_photos
```

This will copy all files from `./photos` (and its subdirectories) to `./organized_photos`, organizing them in a structure like:

```
organized_photos/
├── 2023/
│   ├── 2023-01-06/
│   │   ├── photo1.jpg
│   │   └── document1.jpg
│   └── 2023-12-06/
│       └── photo2.jpg
└── 2024/
    └── 2024-12-24/
        └── christmas.cr3
```

## Building

To build the application:

```bash
go build -o snaparchive main.go
```

## Features Detail

### File Organization
Files are organized based on their modification time into directories following the pattern:
- `YYYY/YYYY-MM-DD/filename`
- Example: `2025/2025-10-05/myfile.txt`

### Conflict Resolution
The application provides three strategies for handling file conflicts:

1. **Skip (default)**: Skip files that already exist in the destination
2. **Overwrite**: Replace existing files with source files
3. **Rename**: Append a number to make the filename unique (`myfile.txt` → `myfile_1.txt` → `myfile_2.txt`, etc.)

### Error Handling
The application includes comprehensive error handling for:
- Invalid command line arguments
- Non-existent source directories
- Permission issues
- File copy failures

## Requirements

- Go 1.16 or later
- Read access to source directory
- Write access to destination directory (or parent directory to create it)