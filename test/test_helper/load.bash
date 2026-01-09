setup() {
    bats_load_library bats-support
    bats_load_library bats-assert
    export PATH="/app/bin/commands:$PATH"
    export CLAUDECODE=1
}

trun() {
    run timeout 5 "$@"
}
