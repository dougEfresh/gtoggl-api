// Copyright 2016 Douglas Chimento.  All rights reserved.

/*
Package gclient provides access to toggl REST API.

Example:
       import "gopkg.in/dougEfresh/gtoggl.v8"
       import "gopkg.in/dougEfresh/toggl-client.v8"

       func main() {
	    thc, err := gtoggl.NewClient("token")
	    ...
	    tc, err := gclient.NewClient(thc)
	    ...
	    client,err := tc.Get(1)
	    if err == nil {
	 	panic(err)
	   }
       }
*/
package gtogglapi
// Copyright 2016 Douglas Chimento.  All rights reserved.

/*

Example:
        import "gopkg.in/dougEfresh/gtoggl.v8"
        import "gopkg.in/dougEfresh/toggl-timeentry.v8"

        func main() {
	    thc, err := gtoggl.NewClient("token")
	    ...
	    tc, err := gtimeentry.NewClient(thc)
	    ...
	    timeentry,err := tc.Get(1)
	    if err == nil {
		panic(err)
	    }
	}
*/
// Copyright 2016 Douglas Chimento.  All rights reserved.

/*
Package gproject provides access to toggl REST API


Example:
       import "gopkg.in/dougEfresh/gtoggl.v8"
       import "gopkg.in/dougEfresh/toggl-project.v8"

       func main() {
	    thc, err := gtoggl.NewClient("token")
	    ...
	    tc, err := gproject.NewClient(thc)
	    ...
	    project,err := tc.Get(1)
	    if err == nil {
	 	panic(err)
	   }
       }
*/
/// Copyright 2016 Douglas Chimento.  All rights reserved.

/*
Package gtest provides test utils for gtoggl
*/
// Copyright 2016 Douglas Chimento.  All rights reserved.

/*
Package gtimeentry provides access to toggl REST API


Example:
       import "gopkg.in/dougEfresh/gtoggl.v8"
       import "gopkg.in/dougEfresh/toggl-timeentry.v8"

       func main() {
	    thc, err := gtoggl.NewClient("token")
	    ...
	    tc, err := gtimeentry.NewClient(thc)
	    ...
	    timeentry,err := tc.Get(1)
	    if err == nil {
	 	panic(err)
	   }
       }
*/
// Copyright 2016 Douglas Chimento.  All rights reserved.

/*
Package guser provides access to toggl REST API


Example:
       import "gopkg.in/dougEfresh/gtoggl.v8"
       import "gopkg.in/dougEfresh/toggl-user.v8"

       func main() {
	    thc, err := gtoggl.NewClient("token")
	    ...
	    tc, err := gtimeentry.NewClient(thc)
	    ...
	    me,err := tc.Get(false)
	    if err == nil {
	 	panic(err)
	   }
       }
*/
// Copyright 2016 Douglas Chimento.  All rights reserved.

/*
Package gworkspace provides access to toggl REST API


Example:
       import "gopkg.in/dougEfresh/gtoggl.v8"
       import "gopkg.in/dougEfresh/toggl-workspace.v8"

       func main() {
	    thc, err := gtoggl.NewClient("token")
	    ...
	    tc, err := gworkspace.NewClient(thc)
	    ...
	    workspace,err := tc.Get(1)
	    if err == nil {
	 	panic(err)
	   }
       }
*/

