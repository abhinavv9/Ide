import os

# Retrieve the code from the environment variable
code = os.getenv("CODE")

# Execute the code
exec(code)
