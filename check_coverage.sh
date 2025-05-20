#!/bin/bash

# Run the tests and generate the coverage profile
go test -coverpkg=./... -coverprofile=profile.cov ./test/...

# Get the coverage data
coverage_data=$(go tool cover -func=profile.cov | awk '{print $1, $3}' | sed 's/%//')

# Set the minimum required coverage percentage 
MIN_COVERAGE=93

# File containing the list of files to skip
SKIP_FILES_FILE="unit_test_coverage_skip_files.txt"

# Read the skip files into an array
IFS=$'\r\n' GLOBIGNORE='*' command eval 'SKIP_FILES=($(cat "$SKIP_FILES_FILE"))'

# Array to store processed files
processed_files=()

# Function to check if a file is in the skip list
is_skipped_file() {
    local file="$1"
    for skip_file in "${SKIP_FILES[@]}"; do
        if [[ "$file" == "$skip_file" ]]; then
            return 0
        fi
    done
    return 1
}

# Function to calculate the average coverage for a file
calculate_average() {
    local file=$1
    local total_percentage=0.00
    local line_count=0
    
    # Process each line of coverage data
    while read -r coverage_file coverage_percentage; do
        # Check if the coverage file matches the desired file
        if [[ "$coverage_file" == *"$file"* ]]; then
            total_percentage=$(awk "BEGIN {printf \"%.2f\", $total_percentage + $coverage_percentage}")
            ((line_count++))
        fi
    done <<< "$coverage_data"
    
    # Calculate average percentage
    if (( line_count > 0 )); then
        average_percentage=$(awk "BEGIN {printf \"%.2f\", $total_percentage / $line_count}")
        echo "$average_percentage"
    else
        echo "0"
    fi
}

should_fail=false

# Process each line of coverage data
while read -r file coverage_percentage; do
    # Remove the suffix starting with ":" from file
    file=${file%%:*}

    # Skip files in the skip list
    if is_skipped_file "$file"; then
        continue
    fi
    
    # Check if the file has already been processed
    if [[ " ${processed_files[@]} " =~ " ${file} " ]]; then
        continue
    fi
    
    # Call calculate_average function and pass 'file' as an argument
    average_percentage=$(calculate_average "$file")
    
    # Check if the coverage is below the minimum required
    if (( $(echo "$average_percentage < $MIN_COVERAGE" | awk '{print ($1 < 93) ? 1 : 0}') )); then
        echo "Coverage check failed for file: $file with $average_percentage%"
        should_fail=true
    fi
    
    # Add the processed file to the list
    processed_files+=("$file")
done <<< "$coverage_data"

# Fail the script if any file's coverage is below the minimum required
if $should_fail; then
    echo "Coverage check failed"
    exit 1
fi

echo "All files have sufficient coverage"
exit 0