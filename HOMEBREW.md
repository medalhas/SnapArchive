# Publishing SnapArchive to Homebrew

This guide explains how to publish SnapArchive to Homebrew so users can install it with `brew install snaparchive`.

## Option 1: Personal Tap (Recommended)

### Step 1: Create a Homebrew Tap Repository

1. Create a new GitHub repository named `homebrew-tools` (must start with `homebrew-`)
2. Clone it locally:
   ```bash
   git clone git@github.com:medalhas/homebrew-tools.git
   cd homebrew-tools
   ```

### Step 2: Create the Formula

1. Copy the `snaparchive.rb` file from this repository to your tap repository
2. Update the formula with the correct SHA256 hash after creating a release

### Step 3: Create a GitHub Release

1. **Tag and release your code:**
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

2. **The GitHub Action will automatically:**
   - Build binaries for multiple platforms
   - Create a GitHub release
   - Generate checksums

3. **Update the formula:**
   - Download the source archive from your release
   - Calculate its SHA256: `shasum -a 256 v1.0.0.tar.gz`
   - Update the `sha256` field in `snaparchive.rb`

### Step 4: Test the Formula

```bash
# Test the formula locally
brew install --build-from-source ./snaparchive.rb

# Test that it works
snaparchive --help
```

### Step 5: Publish Your Tap

```bash
cd homebrew-tools
git add snaparchive.rb
git commit -m "Add snaparchive formula"
git push origin main
```

### Step 6: Users Can Install

Users can now install your package with:
```bash
brew tap medalhas/tools
brew install snaparchive
```

## Option 2: Submit to Official Homebrew (For Popular Tools)

If your tool becomes popular, you can submit it to the official Homebrew repository:

### Requirements for Official Homebrew:
- Tool must be stable and actively maintained
- Must have 30+ forks, 30+ watchers, or 75+ stars
- No duplication of existing functionality
- Must pass all Homebrew tests

### Process:
1. Fork `homebrew/homebrew-core`
2. Add your formula to `Formula/snaparchive.rb`
3. Submit a pull request
4. Address review feedback

## Quick Start (Personal Tap)

1. **Create the release:**
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

2. **Create tap repository:**
   ```bash
   # Create https://github.com/medalhas/homebrew-tools
   git clone git@github.com:medalhas/homebrew-tools.git
   cd homebrew-tools
   cp ../SnapArchive/snaparchive.rb .
   ```

3. **Update SHA256 in formula after release is created**

4. **Publish:**
   ```bash
   git add snaparchive.rb
   git commit -m "Add snaparchive formula"
   git push origin main
   ```

5. **Test installation:**
   ```bash
   brew tap medalhas/tools
   brew install snaparchive
   ```

## Files Created for Homebrew Publishing:

- `.github/workflows/release.yml` - GitHub Action for automated releases
- `.goreleaser.yaml` - Configuration for building cross-platform binaries
- `snaparchive.rb` - Homebrew formula
- This guide in `HOMEBREW.md`

## Next Steps:

1. Commit these files to your repository
2. Create your first release (v1.0.0)
3. Create the homebrew tap repository
4. Update the SHA256 in the formula
5. Test and publish!