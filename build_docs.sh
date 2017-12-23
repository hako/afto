if (! ronn_cmd="$(type -p "ronn")" || [ -z "$ronn_cmd" ]) && (! md2manhtml_cmd="$(type -p "md2man-html")" || [ -z "$md2manhtml_cmd" ]) ; then
	echo "You would need to install ronn and md2man-html to build the docs."
	else
	md2man-roff docs/afto.md > docs/afto.1
fi