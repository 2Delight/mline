# MLine

## Team

### Name

0x22E

### Project

Web editor for reproducible pipelines

### Members

- Ksenia Poliakova
- Oleg Sidorenkov
- Anton Timonin
- Egor Timonin
- Tsvetkova Maria Andreevna

## Requirements

### Summary

Web editor for reproducible pipelines designed for creating, editing, validating specifications with a library of predefined and custom types that follow an ML experiment domain model defined as a semantic network represented in YAML. The editor could be used as an extension in VS Code or as Web App (WPA) served from website.

### Stakeholders

- Developers
- DevOps/SRE-engineers
- Q/A-engineers
- Other tech specialists and managers who interact with CI/CD
- Investors
- Society

### Expected Needs

- Delivary releases
- Rollback releases
- Validate and test releases

### Features

- **Viewing specification** - Users have access to their specifications in the web application and also in VS Code extension. So there are two features that are connected with viewing specifications:
	- Viewing YAML - YAML is a user-friendly data serialization format that is well structured, easy to browse and analyze. In this format it is easy for the human eye to see the details of interest.
	- Syntax Highlighting - this feature allows users to avoid confusion while watching specification file, because it also helps human eye to structure information and make it is easy to read and find some defects in realization.

- **Edit specification** - Project provides users with such important possibility like editing configurations, because this is how different defects may be removed. Some details:
	- Editing YAML - As it was written before, YAML is structured, so for users it is very convenient to edit specifications, that are stored in YAML format.
	- Autocompletion - This option is added to the project to help users to edit specifications, so the system can suggest to continue command or job name, etc, so that users do not have to write full commands.

- **Save specification** - Users need to have an option to save their specification files, not to lose some changes that were made. Systems, having option of saving may have two different features: 
	- Manual saves - Users can choose a particular time for saving by themselves, so their specification files might be saved on the particular stage, that users need.
	- Autosaving - System has special feature that helps users not to lose their changes made in specification files in some unexpected situations, such as the electricity went out. So user doesn't have to save everything manually, everything will be saved automatically.

- **Specification Validation** - It is rather important for system to have validation of specification files, that can save much time for users.
	- Customizable linting - Linter scans user's specification files for areas that require to be removed, so it can save user's time.
	- Static Analysis - This feature helps users to detect some bugs or mismatches in specification files.

- **Code Hints** - includes code highlighting, autocompletion and prompts (e. g. entity name automatic suggestions) and highly customizable linting.

- **Repository Integration** - all the changes can be applied to a GitLab/GitHub repository specification file.

- **Framework Integration** - specification state changes affect the MLOps platform via CI/CD framework integration.

### Constraints

- **Security**
The system must provide secure access to specifications. It will be achived via authentication, authorization and ensuring the security of job runners and other infrastructure components. Moreover, we embrace whitelist approach.

- **Maintainability**
The system should be easy to extend with new features and integrations. It will be achieved with multilayered architectury approach of writing code. Also, every entity will have separate non-intersecting functionality which will make system more flexible.

- **Usability**
The platform must support multiple interfaces: VS Code extension and web UI. It lets user to operate efficiently. The UI must be friendly, so user with low experience of using CI/CD can navigate easily. The will also provide a web-editor with auto-completion hints.

- **Reliability**
  - **Availability**: The service have to provide a high level of SLA (e. g. 99.9%) and be reachable/available as much time as possible.
  - **Recoverability**: The system must be able to restore to a functional state after failures within not longer period than. Regular back ups of configurations, results of pipelines and jobs must be provided, too.

- **Observability**
For maintainers debug logging and tracing must be provided. For users, container logs should be provided

## Architecture

### Draft

![Architecture](diagrams/draft-architecture.drawio.svg)
