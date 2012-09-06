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

package term

/***********************************************
# Libraries / Modules
###
#  These are the libraries / modules
#  necessary to compile this program
##
***********************************************/
/*
import (
  "fmt"
  "os"
)
*/

/***********************************************
#  Sub-Routines / Functions Definitions
###
# This is the actual Body
##
***********************************************/



const (
/*


  Standard British Colour Abbreviations according to:

  http://www.electriciansforums.co.uk/electrical-forum-general-electrical-forum/9388-british-standard-abbreviations.html
  According to BS 7645:1993 (IEC 60757 1983)
  Black ~ BK
  Brown ~ BN
  Red ~ RD
  Orange ~ OG
  Yellow ~ YE
  Green ~ GN
  Blue ~ BU
  Violet ~ VT
  Grey ~ GY
  White ~ WH
  Pink ~ PK
  Turquoise ~ TQ

  confirmed :
  http://www.helukabel.de/pdf/english/technik/X_057_Colour_Abbreviations_according_to_VDE_and_IEC.pdf

  F == Foreground ie Text
  B == Background

  RST == Reset

// Colour codes provided by Petar Maymounkov :: http://groups.google.com/group/golang-nuts/browse_thread/thread/f7d30ab4492f4367
// Thank you !

*/
        RST = "\x1b[0m"
        Bright = "\x1b[1m"
        Dim = "\x1b[2m"
        Underscore = "\x1b[4m"
        Blink = "\x1b[5m"
        Reverse = "\x1b[7m"
        Hidden = "\x1b[8m"
        FBK = "\x1b[30m"
        FRD = "\x1b[31m"
        FGN = "\x1b[32m"
        FYE = "\x1b[33m"
        FBU = "\x1b[34m"
        FMagenta = "\x1b[35m"
        FCyan = "\x1b[36m"
        FWH = "\x1b[37m"
        BBK = "\x1b[40m"
        BRD = "\x1b[41m"
        BGN = "\x1b[42m"
        BYE = "\x1b[43m"
        BBU = "\x1b[44m"
        BMagenta = "\x1b[45m"
        BCyan = "\x1b[46m"
        BWH = "\x1b[47m"
)

/************************************************************************
###                                 EOF                               ###
*************************************************************************/

