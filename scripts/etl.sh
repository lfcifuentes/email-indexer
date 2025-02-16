#!/bin/bash

# Move to the backend directory
cd backend

echo "Downloading the data from the source..."

# Download with consistent filename
curl -L --progress-bar -o enron_mail_20110402.tgz http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz

# Check if download was successful and file exists
if [ $? -eq 0 ] && [ -s enron_mail_20110402.tgz ]; then
    echo "Downloaded successfully!"
    echo "Extracting the data..."
    tar -xzf enron_mail_20110402.tgz
    
    if [ $? -eq 0 ]; then
        echo "Extracted successfully!"
        echo "Cleaning up..."
        rm enron_mail_20110402.tgz
        echo "Done! Dataset is ready in backend/enron_mail_20110402/maildir"
    else
        echo "Error: Extraction failed"
        exit 1
    fi
else
    echo "Error: Download failed or file is empty"
    rm -f enron_mail_20110402.tgz
    exit 1
fi