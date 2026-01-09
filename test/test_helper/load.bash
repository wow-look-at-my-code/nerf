setup() {
    bats_load_library bats-support
    bats_load_library bats-assert
    export CLAUDECODE=1
}

trun() {
    run timeout 5 "$@"
}

# Cross-platform script wrapper to simulate a terminal
# Usage: run_in_pty command [args...]
run_in_pty() {
    if [[ "$(uname)" == "Darwin" ]]; then
        run script -q /dev/null "$@"
    else
        run script -q -c "$*" /dev/null
    fi
}
