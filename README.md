# syslite

Minimal system information extractor written in Go. (Trying my best to keep this lightweight)
**Currently supports only YAML config-based execution.**

```bash
echo "You love Bash commands or script... Until you need 20 system infos, YAML configs and env injection."
echo "Copy-pasting the same mess across repos makes you question your life."
```


## Features

* Extracts system data (based on configured fields)
* Supports flexible output rendering

  * File output (`path`)
  * Console output (`stdout`)
* Easily extendable

---

## Directory Structure

```txt
syslite/
├── config/     # YAML config loader
├── extract/    # Extract system data
├── outputs/    # Output renderer
└── main.go     # Entry point
```

---

## Sample `config.yaml`

```yaml
extract:
  - hostname
  - os
  - distro

output:
  formats:
    json:
      path: "out/system.json"
      stdout: true
```

---

##  Run It

### Step 1: Auto-inject env vars (Your binary)

```yaml
- name: Inject sys info
  run: |
    ./syslite  # injects into $GITHUB_ENV
```

---

###  Step 2: Use injected vars

```yaml
- name: Use kernel info
  run: |
    echo "Kernel version: $KERNEL_VERSION"
    if [[ "$KERNEL_MAJOR" == "6" ]]; then
      echo "Modern Linux Kernel"
    fi
```

---

###  Example of what `syslite` injects

```bash
KERNEL_VERSION=6.8.0-xyz
KERNEL_MAJOR=6
HOSTNAME=ci-runner-07
```

---
Boom. Any job after `syslite` now uses those injected env vars natively. Zero parsing. It can support any workflow (eg.Ansible runtime)

---

## Planned

* Support CLI flags (e.g., `--format`, `--stdout`)
* Auto-detect config location
* Extend extractors (CPU, RAM, etc.)
