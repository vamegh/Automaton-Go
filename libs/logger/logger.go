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
package logger

/***********************************************
#  Libraries / Modules
###
#  These are the libraries / modules
#  necessary to compile this program
##
***********************************************/

import (
  "fmt"
  "automaton/libs/term"
  "time"
  "os"
)

/***********************************************
#  Sub-Routines / Functions Definitions
###
# This is the actual Body
##
***********************************************/

var countD, countE int

func Deblog (message string, loglev byte, caller *string)   {
  t := time.Now().UTC()
  date := t.Format("2006_01_02-15:04:05")
  switch loglev {
    case 1:
      fmt.Println (term.FGN+date,"::",term.FBU,"DEBUG ::",
                    term.FYE,countD,*caller,term.FBU,"::",message,term.RST)
    case 2:
      fmt.Println (term.FGN+date,"::",term.FBU,"DEBUG ::",
                    term.FYE,countD,*caller,term.FBU,"::",message,term.RST)
    case 3:
      fmt.Println (term.FGN+date,"::",term.FBU,"DEBUG ::",
                    term.FYE,countD,*caller,term.FBU,"::",message,term.RST)
  }
  countD++
  countE++
}

func Notelog (message string, caller *string)   {
  t := time.Now().UTC()
  date := t.Format("2006_01_02-15:04:05")
  fmt.Println (term.FGN+date,"::",term.FBU,message,term.RST)
}

func Errlog (message *error, loglev byte, caller *string)   {
  t := time.Now().UTC()
  date := t.Format("2006_01_02-15:04:05")
  switch loglev {
    case 1:
      fmt.Println (term.FGN+date,"::",term.FRD,"ERROR ::",
                    term.FYE,countD,*caller,term.FRD,"::",*message,term.RST)
      os.Exit(1)
    case 2:
      fmt.Println (term.FGN+date,"::",term.FRD,"ERROR ::",
                    term.FYE,countD,*caller,term.FRD,"::",*message,term.RST)
      os.Exit(2)
    case 3:
      fmt.Println (term.FGN+date,"::",term.FRD,"ERROR ::",
                    term.FYE,countD,*caller,term.FRD,"::",*message,term.RST)
      os.Exit(3)
  }
  countD++
  countE++
}

/************************************************************************
###                                 EOF                               ###
*************************************************************************/

