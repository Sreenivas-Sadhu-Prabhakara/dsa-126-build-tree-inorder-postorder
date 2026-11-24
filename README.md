# Build Tree from Inorder + Postorder

> Striver SDE Sheet 126 · Day 19 Binary Tree III — the preorder construction mirrored: walk postorder backwards for roots, so build the right subtree first. Go, stdlib only, tested.

## Problem

Reconstruct the unique binary tree from its inorder and postorder traversals (distinct values).

## Approach

Postorder ends with the root — so read it *backwards* and it deals out roots exactly like preorder does forwards, except the next root encountered belongs to the **right** subtree (reversed postorder is root-right-left). Hence the single structural difference from #125: recurse right before left, and the one backwards-moving cursor stays in sync with no index bookkeeping. Same value→index map for O(1) splits, same round-trip tests. Seeing the two constructions as mirror images — forward cursor/left first vs backward cursor/right first — is the understanding the sheet is after by assigning both.

**Complexity:** O(n) time, O(n) map + O(h) recursion.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
