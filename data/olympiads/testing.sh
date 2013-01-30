#!/bin/sh
#
#=============ABOUT=================================================
#testing.sh
#
#For each csv file in this directory, testing.sh creates a collection
#for testing purposes. For this script to work you should have a 
#working local mongodb without username/password protection.
#
#THIS TOOL IS NOT INTENDED FOR DEPLOYMENT
#
#===================================================================

for file in *.csv; do
  collection=${file%.csv}
  echo "Adding collection ${collection} to database"
  mongoimport --type csv --file "$file" --headerline -c "$collection"
done

echo Collections added.
