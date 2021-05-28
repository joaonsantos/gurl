set -e
# first argument is the number of commits to add
echo "" | cat - changelog.txt > temp && mv temp changelog.txt
git log --pretty="- %s" | head -n $1 | cat - changelog.txt > temp && mv temp changelog.txt
echo "## $(date +%Y-%m-%d)" | cat - changelog.txt > temp && mv temp changelog.txt
