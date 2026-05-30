#!yasl_linux_x64

(! !)

"[-----]"

// (r; g; b) -> void
{


	// \e[48;2;r;g;bm
	"\e[48;2"
	3 #(
		";" +
		(2 ^)
		"" +
		+
	)
	"m "
	+
	=
	print
	print
}print_color

"---\n" print
//	(c) -> u6
{
	0		//	ret
	(# !^)
	(# !^)
	0		//	i
	"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	[#] #(
		(2 !^)							// load i
		(%]#[)					// str[i]
		(# ^)							// load c
		=
		(# !^)							// store c

		 == ? (
			=
			(# 2 #^)
			!
			(2 ^)
			(# 2#!^)

		)
		1 +	//	i++
		(2 ^)							// store i
	)
	!
	!
	(# ^)
	!
	(# ^)
}get_base64

// (u6, u6, u6, u6) -> (u8, u8, u8)
{
	18 <<
	(4 !^)
	12 <<
	(4 !^)
	6 <<
	(4 !^)
	+ + +
	= =
	16 >>
	255 &
	(3 !^)
	8 >>
	255 &
	(3 !^)
	255 &
	(3 !^)
}process_base64

{
	""
	read
	@(
		+
		read
	)
	!
}load_img

{
	[#]
	1
	(= = * <)
	@(
		1 +
		(= = * <)
	)
}sqrt

load_img
[#]
0
(2 ^)
#(
		%]#[
		get_base64
		(# !^)
		( = 4 % 3 ==)?
	(
		(# 4 #^)
		process_base64
		print_color
	):(
	)


	1 +
)
