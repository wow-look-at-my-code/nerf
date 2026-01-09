#!/usr/bin/env bats

bats_load_library test_helper

@test "touch blocks modifying existing files" {
    trun touch test/fixtures/sample.txt
    assert_failure 1
    assert_output --partial "not allowed on existing files"
}

@test "touch allows creating new files" {
    local tmpfile="$BATS_TEST_TMPDIR/newfile.txt"
    trun touch "$tmpfile"
    assert_success
    [ -f "$tmpfile" ]
}
