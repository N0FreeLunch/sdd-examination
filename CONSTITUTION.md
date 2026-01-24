# Project Constitution

This document defines the **absolute rules** that govern the development and maintenance of this project. All contributors (humans and AI) must adhere to these rules without exception.

## Article I. Separation of Concerns
1.  **External Specifications**: The requirements, domain logic, and API designs for this project are strictly defined in an **external repository**.
2.  **No Spec Commits**: You **MUST NEVER** commit the contents of the `sdd-examination-spec` folder (or any other path containing the specs) to this repository. The specification repository must remain a separate entity to prevent data duplication and versioning conflicts.
3.  **Git Ignore & Verification**: 
    *   The symbolic link `sdd-examination-spec` must always be present in `.gitignore`.
    *   **Pre-Commit Check**: Before running `git add .`, you MUST check `git status` to ensure `sdd-examination-spec` is NOT being tracked or staged. If it appears as `modified` or `new file`, you MUST untrack it (`git rm --cached sdd-examination-spec`) immediately.

## Article II. Specification-Driven Development (SDD)
1.  **Spec First**: Implementation code shall not be written until a corresponding specification exists.
2.  **Single Source of Truth**: When ambiguity arises in the code, the Specification is the final authority. Do not change the code to fit assumptions; verify the Spec first.
3.  **Verification**: Code changes must be verified against the rules defined in the Specification.

## Article III. AI Assistant Guidelines
1.  **Read-Only Specs**: AI assistants may read `sdd-examination-spec` to understand requirements but must **never** attempt to modify files within that directory unless explicitly instructed to update the *Specification Repository* itself.
2.  **Cross-Language Implementation**: While specifications may be written in Korean (or other languages), the implementation code (variables, comments, commit messages) must be written in **English**.

## Article IV. Amendments
1.  **User Confirmation Required**: Any changes to this **CONSTITUTION.md** file require explicit approval from the USER. AI assistants must not modify this file without a direct request or confirmation.
