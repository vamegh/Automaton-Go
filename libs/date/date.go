/*************************************************************************
#                                                                        #
#       Automaton Framework re-implemented in Golang                     #
#                                                                        #
#       Automation & Auditing Toolkit                                    #
#                                                                        #
#       Copyright (C) 2010 by Vamegh Hedayati                            #
#       Golang re-implementation Copyright (C) 2012 by Vamegh Hedayati   #
#       vamegh@gmail.com                                                 #
#                                                                        #
#       Please see Copying for License Information                       #
#                               GNU/GPL v2 2010 / 2012                   #
*************************************************************************/

/*
       Automation & Auditing Toolkit

       Copyright (C) 2010/2012 by Vamegh Hedayati

       Please see Copying for License Information
                   GNU/GPL v2 2010 / 2012
*/
package date

/***********************************************
# Libraries / Modules
###
#  These are the libraries / modules
#  necessary to compile this program
##
***********************************************/
import (
  "time"
)

/***********************************************
#  Sub-Routines / Functions Definitions
###
# This is the actual Body
##
***********************************************/
func Make_date() (time.Time) {
  current_time := time.Now()
  /*
    Month := time.Month()
    Second := time.Second
    format := time.Time.Format("dd-mm-yyyy hh:mm:ss");
    fmt.Println("Using a default format", Second)
  */
  return current_time
}



/************************************************************************
###                                 EOF                               ###
*************************************************************************/

