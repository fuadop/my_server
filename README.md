Example usage of sendchamp go lib in an http server.

### Note
In your dotenv file.. define your public key with single quotes around the value, bash assumes "`$`" as variables so they might get replaced or ignored.

```bash
SENDCHAMP_PUBLIC_KEY='sendchamp_test_$2y$10$U2SHG5T2F/cr0jfzNCKgguHv.23plvJP/75EzZjF5MtLXz65SDrQi'
```

If the single quotes in the above was ignored, bash would replace `$2`, `$1` and `$U` with blank values making your token invalid.
