#!/bin/sh

extension="$1"
project_dir=$(realpath "$(dirname "$(readlink -f "$0")")/../")
binaries_dir="$project_dir/.bin"
file_prefix="etcd-tui_$GOOS-$GOARCH"

if [ -n "$extension" ]
then
binary_file_name="$file_prefix.$extension"
else
binary_file_name="$file_prefix"
fi

binary_path="$binaries_dir/$binary_file_name"
archive_path="$binaries_dir/$file_prefix.tar.gz"

rm -f "$archive_path"
rm -f "$binary_path"

(cd "$project_dir" && go build -o "$binary_path") || {
    echo "build failed"
    exit 1
}

(cd "$binaries_dir" && tar -czvf "$archive_path" "$binary_file_name") || {
    rm -f "$binary_path"
    echo "archive failed"
    exit 1
}

echo "archive created $archive_path"
