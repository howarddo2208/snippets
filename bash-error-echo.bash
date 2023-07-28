# echo the error with color red and append to standard error
echoerr() { echo -e "\033[31m$@\033[0m" 1>&2; }

# echoerr this is an error
