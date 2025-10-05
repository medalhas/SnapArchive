# SnapArchive

An application that recursively reads files from a source directory and copies them to a destination directory organized by year and month based on the file's modification time.
Mainly to organize my photos, but can work on any type of filesg

## Features

- Recursively processes all files in the source directory and its subdirectories
- Organizes files by year and month (YYYY/MM-MonthName format)
- Handles file name conflicts by appending a number to duplicate names
- Preserves file permissions
- Provides clear error messages and validation
- Shows progress during file copying

## Usage

```bash
./snaparchive <source_folder> <dest_folder>
```

### Arguments

- `source_folder`: The path to the directory containing files to organize
- `dest_folder`: The path where organized files will be copied

### Example

```bash
./snaparchive ./photos ./organized_photos
```

This will copy all files from `./photos` (and its subdirectories) to `./organized_photos`, organizing them in a structure like:

```
organized_photos/
├── 2023/
│   ├── 01/
│   │   ├── photo1.jpg
│   │   └── document1.jpg
│   └── 12-December/
│       └── photo2.jpg
└── 2024/
    └── 06/
        └── file1.cr3
```

## Building

To build the application:

```bash
go build -o snaparchive main.go
```

## Features Detail

### File Organization
Files are organized based on their modification time into directories following the pattern:
- `YYYY/MM-MonthName/filename`
- Example: `2024/10-October/myfile.txt`

### Conflict Resolution
If a file with the same name already exists in the destination, the application automatically appends a number to make it unique:
- `myfile.txt` → `myfile_1.txt` → `myfile_2.txt`, etc.

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