## Github action

### Reference

[Github Action Marketplace](https://github.com/marketplace?type=actions)

### How to

- create a folder name .gitub in root folder
- create file ci.yaml

sample ci.yaml file for .net:

```yaml
name: Continuous Integration
on: push # trigger for the workflow (here it is push event)
jobs:
  run_code_checks:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Repo Code
        uses: actions/checkout@v2
      - name: Setup .Net
        uses: actions/setup-dotnet@v3
        with:
          dotnet-version: "6.0.x"
          include-prerelease: true
      - name: Install dependencies
        run: dotnet restore
      - name: Build
        run: dotnet build --configuration Release --no-restore
      - name: Test
        run: dotnet test --no-restore --verbosity normal
```
