# exec_out style plugin for fluent-bit

# Example Configuration

```
[PLUGINS]
    Path /path/to/out_exec.so

[OUTPUT]
    Name exec_out
    Match nginx.*
    Command python3 /path/to/script.py
```

# References

- fluentd exec_out plugin documentation <https://docs.fluentd.org/output/exec>
- go output plugin documentation <https://docs.fluentbit.io/manual/development/golang-output-plugins>
- Other Example plugins <https://github.com/fluent/fluent-bit-go/network/dependents>
