#!/usr/bin/env bats

bats_load_library test_helper

@test "find -delete is blocked" {
    trun find /tmp -name "*.tmp" -delete
    assert_failure
    assert_output --partial "not allowed"
}

@test "find -exec is blocked" {
    trun find /tmp -exec ls {} \;
    assert_failure
    assert_output --partial "not allowed"
}

@test "find with -name uses fd" {
    trun find /etc -name "passwd" -maxdepth 1
    assert_success
    assert_output --partial "passwd"
}

@test "find -name with glob pattern works" {
    mkdir -p /tmp/nerf-test
    touch /tmp/nerf-test/test.go /tmp/nerf-test/test.txt
    trun find /tmp/nerf-test -name "*.go"
    assert_success
    assert_output --partial "test.go"
    refute_output --partial "test.txt"
    rm -rf /tmp/nerf-test
}

@test "find -name with wildcard in middle works" {
    mkdir -p /tmp/nerf-test
    touch /tmp/nerf-test/foo_bar.txt /tmp/nerf-test/foo_baz.txt /tmp/nerf-test/other.txt
    trun find /tmp/nerf-test -name "foo_*.txt"
    assert_success
    assert_output --partial "foo_bar.txt"
    assert_output --partial "foo_baz.txt"
    refute_output --partial "other.txt"
    rm -rf /tmp/nerf-test
}
