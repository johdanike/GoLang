# find . -name "*sh" | sed 's|.*/||' | sort -r
find . -name "*.sh" | xargs -n 1 basename -s .sh | sort -r
