default: generate

generate:
    controller-gen paths=. rbac:roleName=constrained output:rbac:dir=./rbac