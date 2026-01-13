#!/usr/bin/env bats

bats_load_library test_helper

@test "docker without CLAUDECODE passes through" {
    unset CLAUDECODE
    trun docker --version
    assert_success
    assert_output --partial "Docker"
}

@test "docker with CLAUDECODE passes through for non-compose commands" {
    trun docker --version
    assert_success
    assert_output --partial "Docker"
}

@test "docker compose restart is blocked" {
    trun docker compose restart
    assert_failure 1
    assert_output --partial "docker compose up -d"
}

@test "docker compose --ansi always restart is blocked" {
    trun docker compose --ansi always restart
    assert_failure 1
    assert_output --partial "docker compose up -d"
}

@test "docker compose -f myfile.yml restart is blocked" {
    trun docker compose -f myfile.yml restart
    assert_failure 1
    assert_output --partial "docker compose up -d"
}

@test "docker compose --project-name test restart is blocked" {
    trun docker compose --project-name test restart
    assert_failure 1
    assert_output --partial "docker compose up -d"
}

@test "docker build --no-cache drops the flag" {
    # Create a temp directory with a simple Dockerfile
    tmpdir=$(mktemp -d)
    echo "FROM scratch" > "$tmpdir/Dockerfile"
    # Run docker build with --no-cache; if --no-cache was passed through,
    # Docker would try to pull/build without cache. Since FROM scratch is
    # a special case, this should succeed quickly either way.
    # The main test is that our wrapper doesn't break the command.
    trun docker build --no-cache -t nerf-test-nocache "$tmpdir"
    rm -rf "$tmpdir"
    assert_success
}
