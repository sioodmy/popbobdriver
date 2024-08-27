{buildGoModule}:
buildGoModule {
  pname = "popbobdriver";
  version = "0.0.1";

  src = ./.;

  vendorHash = "sha256-nvVl0/FpSpR73sk3eO8CMNZyX3XrTMnJ+uMa4/57HB0=";

  ldflags = ["-s" "-w"];
}
