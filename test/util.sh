#! /bin/bash

# Output start message of a check
#
# @param string $1 check name
function test_start()
{
	printf "\n\e[1;33m%-40s" "• Running $1"
}

# Output a nice [✔] with colors for a successfully check
function test_success()
{
	echo -en "  \e[1;33m[\e[1;32m✔\e[1;33m]\e[0m"
}

# Output a nice [✘] with colors for a unsuccessfully check
function test_error()
{
	echo -en "  \e[1;33m[\e[1;31m✘\e[1;33m]\e[0m"
}

# Output a nice [✘] with colors for a unsuccessfully check
function test_warning()
{
	echo -en "  \e[1;33m[\e[1;36m!\e[1;33m]\e[0m"
}

function test_end()
{
	echo -en "\e[0m"
}

# @param string $1 output of the command
function print_output()
{
	printf "%s\n" "$(echo -e "$1" | awk '{printf "\t%s\n", $0}')"
}

# @param string $1 output of the command
function print_error()
{
	printf "%s\n" "$(echo -e "$1" | awk '{printf "\t%s\n", $0}')"
}

# @param string $1 output of the command
function print_title()
{
	printf "\033[1m%s\033[0m\n" "$(echo -e "$1" | awk '{printf "\t%s\n", $0}')"
}
