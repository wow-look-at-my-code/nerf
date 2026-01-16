#!/usr/bin/env bats

bats_load_library test_helper

@test "bash script can use sed on files (CLAUDECODE disabled in scripts)" {
    # Create a test script that uses sed on a file
    tmpdir=$(mktemp -d)
    echo "original content" > "$tmpdir/test.txt"
    cat > "$tmpdir/script.sh" << 'EOF'
#!/bin/bash
sed -i 's/original/modified/' "$1"
EOF
    chmod +x "$tmpdir/script.sh"

    # This should work because bash unsets CLAUDECODE
    trun bash "$tmpdir/script.sh" "$tmpdir/test.txt"
    assert_success

    # Verify the file was modified
    run cat "$tmpdir/test.txt"
    assert_output "modified content"

    rm -rf "$tmpdir"
}

@test "bash -c works normally" {
    trun bash -c 'echo hello'
    assert_success
    assert_output "hello"
}

@test "bash interactive mode works" {
    trun bash -c 'echo $((1+2))'
    assert_success
    assert_output "3"
}

@test "bash script can use find with -exec (CLAUDECODE disabled)" {
    tmpdir=$(mktemp -d)
    touch "$tmpdir/a.txt" "$tmpdir/b.txt"
    cat > "$tmpdir/script.sh" << 'EOF'
#!/bin/bash
find "$1" -name "*.txt" -exec basename {} \; | sort
EOF
    chmod +x "$tmpdir/script.sh"

    trun bash "$tmpdir/script.sh" "$tmpdir"
    assert_success
    assert_line "a.txt"
    assert_line "b.txt"

    rm -rf "$tmpdir"
}
