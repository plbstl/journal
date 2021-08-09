# How To Use

notes server - spins up web server on localhost:[port]
	port can be gotten from config or cli flag

notes version

notes help

notes new - alias for "notes create note --open"

notes settings - show configured options (~/.config/notes/config.yml)
	--edit -e - open config in a text editor set in config
	--add - add a config. eg. --add server.port 9001
	--remove --rm - remove a config. eg. --remove server.port
	--all -a - shows all available options and defaults

notes launch - launch the terminal app
	--web - spins up web server on localhost:[port]
			port can be gotten from config or cli flag (--web 9001)
notes create - creates a new note with a default custom id
	*note - --text --open
	*author - creates a new author
	*tag - creates a new tag
	--id - set a custom id to associate with note,author,tag, pls donot use duplicate IDs
notes open <id> - open a note with the specified ID
notes delete [id] - delete a note with the custom id (echo err if no cid found)
	--all - deletes all authors,tags or notes. when deleting authors,tags, notes are not deleted
	(delete --all authors)
notes list - print out a list of created notes (trucated; title, body, 2 lines each)
	--authors - print out a list of authors
	--tags - print out a list of created tags
	--limit - max number of items to print out (default 10)
notes show [id] - print out the contents of a note with the custom id (echo err if no cid found)
	--live - show in terminal app
notes search [flag] [keyword]
	--title (2)
	--author (0)
	--tag (1)
	search order is 0,1,2
notes edit <id> - edit an author,tag

notes import/export [flag] [path]
	--notes
	--authors
	--tags

<!-- notes init - start a new git repo for tracking notes -->
<!-- notes save - commits current changes -->
	<!-- --message --mssg --msg -m -->

<!-- settings -->
server.port = 9001
notes.editor = in-built // or path to editor
cmd.new = open // flags


<!-- ss -->
show last mod date n author
