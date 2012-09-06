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
package filehandler

/***********************************************
#  Libraries / Modules
###
#  These are the libraries / modules
#  necessary to compile this program
##
***********************************************/
import (
  "io/ioutil"
  "regexp"
  "strings"
  "automaton/libs/logger"
  "fmt"
)

/***********************************************
#  Sub-Routines / Functions Definitions
###
# This is the actual Body
##
***********************************************/
func read_file (cfg_file *string) (filein []string){
  caller := "filehandler :: read_file"
  fil,err := ioutil.ReadFile(*cfg_file)
  if err != nil {
    logger.Errlog (&err,1,&caller)
  }
  filein = strings.Split(string(fil), "\n")
  return filein
}

//func compilePattern (pattern

func Read_cfg (cfg_type string, cfg_file *string, cfg_handle string) {
  caller := "filehandler :: read_cfg"
  switch cfg_type {
    case "cfgpars":
      fcon := read_file(cfg_file)
      //patmatch := "(\\w+)=(\".+?\"|\'.+?\'|.+)\\s*(#.*|)"
      // patmatch := "(\\w+)[=](\\w+)"
      wshmatch := "[;]*[#]*(\\s)*[#]."
      varmatch := "(\\w+)=(\".+?\"|'.+?'|.+)(\\s*)(#.*|)"
      wshpat,err := regexp.Compile(wshmatch)
      if err != nil {
        logger.Errlog (&err,1,&caller)
      }
      varpat,err := regexp.Compile(varmatch)
      if err != nil {
        logger.Errlog (&err,1,&caller)
      }

      fmt.Println(wshpat.String())
      fmt.Println(varpat.String())

      for i := 0;i<len(fcon);i++ {
        switch {
          case varpat.MatchString(fcon[i]):
            logger.Deblog("Printing and saving: "+fcon[i],1,&caller)
            variables := strings.SplitAfter(fcon[i], "=")
            for i := 0; i < len(variables);i++ {
              logger.Deblog("Variable currently: "+variables[i],1,&caller)
            }
          case wshpat.MatchString(fcon[i]):
            // skipping the matched pattern
          default:
            logger.Deblog("Hit Default value currently: "+fcon[i],1,&caller)
          }
      }


      /*
          if match,err := regexp.MatchString("^[#]+",fcon[i]); match {
            if err != nil {
              logger.Errlog (&err,1,&caller)
            }
            logger.Deblog("skipping the matched line: "+fcon[i],1,&caller)
          } else if match,err := regexp.MatchString("^(\\s)+",fcon[i]); match {
            if err != nil {
              logger.Errlog (&err,1,&caller)
            }
            logger.Deblog("skipping the matched line: "+fcon[i],1,&caller)
          } else if match := pattern.MatchString(fcon[i]); match {
              if err != nil {
                logger.Errlog (&err,1,&caller)
              }
            logger.Deblog("skipping the patmatch matched line: "+fcon[i],1,&caller)
          } else if match,err := regexp.MatchString("^$",fcon[i]); match {
            if err != nil {
              logger.Errlog (&err,1,&caller)
            }
            logger.Deblog("skipping the matched line: "+fcon[i],1,&caller)
          } else {
            logger.Deblog("Printing and saving: "+fcon[i],1,&caller)
            variables := strings.SplitAfter(fcon[i], "=")
            for i := 0; i < len(variables);i++ {
              logger.Deblog("Variable currently: "+variables[i],1,&caller)
            }
          }
      */
    case "genericpars":
    case "rulespars":
    case "passpars":
    case "shadpars":
    case "grouppars":
    case "userpars":
    case "mailpars":
      switch cfg_handle {
        case "generic":
        case "alert":
        default:
      }
    default:
  }
}


/************************************************************************
###                                 EOF                               ###
*************************************************************************/

