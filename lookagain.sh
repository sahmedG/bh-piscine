find . -type f -name "*.sh" | sort -nr | cut -f 2 -d '.' | sed 's/^\///;s/\// /g'