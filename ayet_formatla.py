import re
import sys

# Read the content of the file
file_path = sys.argv[1]
with open(file_path, 'r', encoding='utf-8') as file:
    content = file.read()

# Define the regex pattern to find the verse format
pattern = r'(\s+\d{1,3}:\d{1,3})'

# Replace the found patterns with the desired markdown format
def replace_with_markdown(match):
    sure_ayet = match.group(0).replace(" ", "").split(':')
    return f" [{sure_ayet[0]}:{sure_ayet[1]}](https://acikkuran.com/{sure_ayet[0]}/{sure_ayet[1]})"

new_content = re.sub(pattern, replace_with_markdown, content)

# Write the new content back to the file
with open("out.md", 'w', encoding='utf-8') as file:
    file.write(new_content)
