@lldnl = internal constant [6 x i8] c"%lld\0A\00"

@runtime_error  = internal constant [14 x i8] c"runtime error\00"

declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)
declare i32 @puts(i8*)
declare i64 @getline(i8**, i64*, %struct._IO_FILE*)

declare void @exit(i32)

declare void @free(i8*)
declare i8* @calloc(i64, i64)
declare i8* @malloc(i64)

declare void @memset(i8*, i8, i64)

declare i64 @strlen(i8*)
declare i8* @strcpy(i8*, i8*)
declare i8* @strcat(i8*, i8*)

%struct._IO_FILE = type opaque

@stdin = external global %struct._IO_FILE*, align 8

define void @printInt(i64 %x) {
       %t0 = getelementptr [6 x i8], [6 x i8]* @lldnl, i32 0, i32 0
       call i32 (i8*, ...) @printf(i8* %t0, i64 %x)
       ret void
}

define void @printString(i8* %s) {
entry:  call i32 @puts(i8* %s)
	ret void
}

define i64 @readInt() {
entry:	%res = alloca i64
        %t1 = getelementptr [6 x i8], [6 x i8]* @lldnl, i32 0, i32 0
	%scanf_out = call i32 (i8*, ...) @scanf(i8* %t1, i64* %res)
	%error = icmp slt i32 %scanf_out, 0
    br i1 %error, label %error_label, label %return_label

error_label:                                                ; preds = %0
    	call void @error()
        unreachable

return_label:                                               ; pred = %0, error_label
	%t2 = load i64, i64* %res
	ret i64 %t2
}

define i8* @readString() {
	%str_ptr = alloca i8*, align 8
	%n = alloca i64, align 8
	store i8* null, i8** %str_ptr, align 8
	store i64 0, i64* %n, align 8
	%stdin = load %struct._IO_FILE*, %struct._IO_FILE** @stdin, align 8
	%str_size = call i64 @getline(i8** %str_ptr, i64* %n, %struct._IO_FILE* %stdin)
	%error = icmp slt i64 %str_size, 0
	br i1 %error, label %error_label, label %return_label

error_label:                                                ; preds = %0
	call void @error()
    unreachable

return_label:                                               ; pred = %0, error_label
	%str = load i8*, i8** %str_ptr, align 8
    %new_line = add i64 %str_size, -1
    %new_line_ptr = getelementptr i8, i8* %str, i64 %new_line
    store i8 0, i8* %new_line_ptr, align 1
	ret i8* %str
}

define void @error() {
	%error_msg = getelementptr [14 x i8], [14 x i8]* @runtime_error, i32 0, i32 0
	call i32 @puts(i8* %error_msg)
	call void @exit(i32 1)
	unreachable
}

define i8* @.str_add(i8* %str, i8* nocapture readonly %str2) {
  	%len = call i64 @strlen(i8* %str)
  	%len2 = call i64 @strlen(i8* %str2)
  	%lenH = add i64 %len, 1
  	%lenSum = add i64 %lenH, %len2
  	%new = call i8* @malloc(i64 %lenSum)
  	%error = icmp eq i8* %new, null
	br i1 %error, label %error_label, label %return_label

error_label:                                                ; preds = %0
	call void @error()
    unreachable

return_label:     											; preds = %0
  	call i8* @strcpy(i8* %new, i8* %str)
  	call i8* @strcat(i8* %new, i8* %str2)
  	ret i8* %new
}

define i8* @.new_arr(i64 %nmemb, i64 %size) {
    %arr_size = mul i64 %nmemb, %size
    %total_size = add i64 %arr_size, 8
    %arr = call i8* @malloc(i64 %total_size)
    call void @memset(i8* %arr, i8 0, i64 %total_size)
    %nmemb_ptr = bitcast i8* %arr to i64*
    store i64 %nmemb, i64* %nmemb_ptr
    %arr_start = getelementptr i8, i8* %arr, i64 8
    ret i8* %arr_start
}