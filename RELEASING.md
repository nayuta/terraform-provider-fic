# Releasing

The release workflow builds provider archives, includes the Terraform Registry
manifest in the checksums, signs the checksum file with GPG, and publishes the
assets to a GitHub release when a `v*` tag is pushed.

Before the first release:

1. Create or select a dedicated GPG signing key and securely back up its private
   key. Terraform Registry accepts RSA and DSA keys.
2. Add the ASCII-armored private key as the `GPG_PRIVATE_KEY` GitHub Actions
   repository secret and its passphrase as `PASSPHRASE`.
3. If publishing on the Terraform Registry, add the ASCII-armored public key to
   the signing keys for the appropriate namespace, and register the provider.
   See [HashiCorp's publishing guide](https://developer.hashicorp.com/terraform/registry/providers/publishing).

To release, check the intended commit and version, then push an annotated tag
such as `v0.5.7` to the fork. Monitor the release workflow and verify the
published archives, manifest, checksums, and detached signature. Registering
and publishing the provider on the Terraform Registry are separate steps.
