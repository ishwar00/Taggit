# Taggit
A tool to help users organize their files using tags.
We proposed this tool as a solution to the problem statement provided to us in the Tally Codebrewers 2022.

# Steps to Setup
<ol>
<li> Clone our repository. You can download the zip file by clicking <a href="https://github.com/singhankit62000/Taggit/archive/refs/heads/master.zip">here</a>.
<li> Install Go from <a href="https://go.dev/dl/">here</a> to compile and run the code and obtain exe files for the tool.
<li> Add 2 items in the context menu namely <strong>Manage Tags</strong> and <strong>Search By Tag</strong>. You can name your context menu items anything you like.
  
<br>

In the Registration editor(regedit.exe), find:
<br>

  ``` 
  1. HKEY_CLASSES_ROOT\Directory\Background\shell if you are administrator
  2. HKEY_CURRENT_USER\Software\Classes\directory\Background\shell if you are a normal user
  ```

<ul>
<li> Add a new key under shell, naming it <strong>Search By Tag</strong>
<li> Add a new key inside this key, named command (mandatory name)
<li> Edit the default property in command to the path link of the <strong>/searchTags.exe "%V"</strong> to pass the file path and name of the selected file to your custom program <br>
</ul>

Also in the Registration editor(regedit.exe), find:
<br>

  ``` 
  1. HKEY_CLASSES_ROOT\*\shell if you are administrator
  2.  HKEY_CURRENT_USER\Software\Classes\*\shell if you are a normal user
  ```

<ul>
<li> Add a new key under shell, naming it <strong>Manage Tags</strong>
<li> Add a new key inside this key, named command (mandatory name)
<li> Edit the default property in command to the path link of the <strong>/manageTags.exe "%1"</strong> to pass the file path and name of the selected file to your custom program <br>
</ul>

<li> Now, whenever you right click on a file, there will be an option called <strong>"Manage Tags"</strong>, clicking on that option enables you to manage tags associated to that file.
<li> Also, another option named <strong>"Search By Tag"</strong> enables the searching feature by taking a tag as an input. It displays all the files that are associated with that tag.
</ol>

# Contributors 

1. <a href="https://github.com/ishwar00">Ishwar</a>
2. <a href="https://github.com/singhankit62000/">Ankit Singh</a>
