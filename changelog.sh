set -e

# set the name of the changelog file
changelog=CHANGELOG.md

# first argument is the number of commits to add
echo "" | cat - $changelog > temp && mv temp $changelog
git log --pretty="- %s" | head -n $1 | cat - $changelog > temp && mv temp $changelog
echo "## $(date +%Y-%m-%d)" | cat - $changelog > temp && mv temp $changelog
