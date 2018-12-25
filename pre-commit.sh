#!/bin/bash

# GOAL : error (go vet) and lint (gofmt) checks using golang vanilla install 

# Inspired from :
# - https://golang.org/misc/git/pre-commit?m=text
# - https://ngurajeka.com/post/using-pre-commit-on-golang-project/

STAGED_GO_FILES=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')

if [[ "$STAGED_GO_FILES" = "" ]]; then
  exit 0
fi

PASS=true

# Run gofmt on all staged files and check the exit status
unformatted=$(gofmt -l $STAGED_GO_FILES)
if [ -n "$unformatted" ]; then
   # Some files are not gofmt'd. Print message and fail.

   echo >&2 "Go files must be formatted with gofmt. Please run:"
   for fn in $unformatted; do
	echo >&2 "  gofmt -w $PWD/$fn"
   done
fi

for FILE in $STAGED_GO_FILES
do

  # Run govet on the staged file and check the exit status	
  go vet $FILE
  if [[ $? != 0 ]]; then
    printf "\t\033[31mgo vet $FILE\033[0m \033[0;30m\033[41mFAILURE!\033[0m\n"
    PASS=false
  else
    printf "\t\033[32mgo vet $FILE\033[0m \033[0;30m\033[42mpass\033[0m\n"
  fi
done

if ! $PASS; then
  printf "COMMIT FAILED\n"
  exit 1
else
  printf "COMMIT SUCCEEDED\n"
fi

exit 0
