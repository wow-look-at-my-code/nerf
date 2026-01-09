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
