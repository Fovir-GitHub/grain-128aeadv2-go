{
  description = "Devshell for Golang.";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux";
    pkgs = import nixpkgs {inherit system;};
  in {
    devShells.${system}.default = pkgs.mkShell {
      # Add packages here.
      buildInputs = with pkgs; [
        esbuild
        go
        golangci-lint
        gopls
        gotools
        govulncheck
        just
        prettier
        tygo
        typescript
      ];

      # Shell hooks.
      shellHook = ''
        echo "Entering the development environment!"
        go version
      '';
    };
  };
}
