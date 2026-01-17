setup() {
    bats_load_library bats-support
    bats_load_library bats-assert
    export CLAUDECODE=1
    unset NERF_IN_SCRIPT
}

# Run command through fake claude to create proper process ancestry
# Process tree: command → bash → claude
# This lets ShouldWrap() detect we're running from "Claude"
trun() {
    run timeout 5 claude "$@"
}

# Cross-platform script wrapper to simulate a terminal
# Runs through fake claude to create proper process ancestry
# Usage: run_in_pty command [args...]
run_in_pty() {
    if [[ "$(uname)" == "Darwin" ]]; then
        run script -q /dev/null claude "$@"
    else
        run script -q -c "claude $*" /dev/null
    fi
}
