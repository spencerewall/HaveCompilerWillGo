## Package Naming
- By convention, packages are given lower case, single-word names; there should be no need for underscores or mixedCaps. 
- Err on the side of brevity, since everyone using your package will be typing that name
- don't worry about collisions a priori
- the package name is the base name of its source directory; the package in src/encoding/base64 is imported as "encoding/base64" but has name base64, not encoding_base64 and not encodingBase64

### References
- https://go.dev/blog/package-names
- https://go.dev/doc/effective_go#package-names