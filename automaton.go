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

       vamegh@gmail.com

       Please see Copying for License Information
                   GNU/GPL v2 2010 / 2012
*/
package main


/***********************************************
# Libraries / Modules
##
#  These are the libraries / modules
#  necessary to compile this program
##
***********************************************/
import (
  //"fmt"
  //"io"
  //"crypto/sha512"
  log "automaton/libs/logger"
  "automaton/libs/filehandler"
)

/***********************************************
#  Sub-Routines / Functions Definitions
##
# This is the actual Body
##
***********************************************/
func main () {
  // This is the main program which calls in everything else.
  caller := "Automaton :: Main ::"

  log.Notelog("Welcome to Automaton",&caller)
  log.Deblog("Development Starting oh yeah !!",3,&caller)

  /*
    it := "Hello, Hello"
    hit := sha512.New()
    io.WriteString(hit,it)
    fmt.Printf("% x", hit.Sum(nil))
    log.DebugL("hello", "debug", 1, "main")
    log.DebugL("hello", "debug", 3, "main")
    log.DebugL("Oh Noes ... ", "error", 3, "main")
  */

//func It (message string, logtype string, loglev int, caller string)
//func Read_cfg (cfg_type *string, cfg_file *string, cfg_handle *string)

  filen := "/home/vdopey/example.txt"
  filet := "cfgpars"
  fileh := ""
  filehandler.Read_cfg (filet,&filen,fileh)

}


/************************************************************************
 ###                                 EOF                              ###
*************************************************************************/

