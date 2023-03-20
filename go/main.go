package main

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)

func main() {
	in := `
    <h1> Git Guide <img src="https://raw.githubusercontent.com/FortAwesome/Font-Awesome/6.x/svgs/brands/git-alt.svg" width="25" height="25">
    </h1>
    
    ## Overview
    
    > Git is a version control system used for tracking changes made to files over time. <br>
    > It helps organize code in a central place, promoting and supporting collaboration on different projects.
    
    > This is incredibly useful for teams working together on projects and requires some thought and heeding good practices to minimize any pain and to best promote efficient, fun, and smooth collaborating. 
    
    <br>
    
    With git, it takes at least three separate commands to share a commit with your colleagues:
    
    The first step adds all files with changes from the working copy of code on your machine to the staging area:
    
    shell 
    $ git add .
    
    
    From here, these changes can be committed to your local repository.
    
    We do this in the second step, passing -m command-line flag to specify we are providing a message describing the changes of the commit, inline, with this commit command:
    
    shell
    $ git commit -m "message to associate with commit"
    
    
    Finally we simply run git push to push the changes on our local repository up to the remote repository for others on your team to pull down on their side. 
    
    shell
    $ git push
    
    
    Notice that we could have used * or . with the first command. Each command results in different behaviors and we should consider the differences now. 
        
    In a nutshell, git add . adds all files, folders, and subfolders, including .gitignore and all other hidden or dot files. 
    
    git add * does virtually the same thing, with the one difference being it will not add any dot files (files beginning with a dot . -- 
        hidden files.
    
    You likely won't run into many cases where you wish to ignore changes made to hidden or dot files, but using git add * would be the command to push any other file changes.
    
    With that said, you will principally use git add . to add local file changes to the 
    staging area.
    
    Although not particularly important, it should be noted that * is
    expanded for you, while . is a git-specific argument that the Git cli uses to add all changes.
    
    ## About This Guide
    ---
    This guide intentionally assumes you have minimal working knowledge at best. 
    
    Each section is organized in a way to be as conducive toward your collaborating with your team and the shared codebase(s) as possible
    
    The Common Workflow section is the core of this guide, with subsequent sections focusing on less commonly used yet useful flows.
    
    The words workflow, and  flow are used often in this guide and are interchangeable. It simply refers to the commands and actions you must take to achieve some common end result desired when working on a team using Git.
    
    ## Branching Strategy/Model
    ---
    
    <br>
    
    *There was once a developer in Helsinki, who accidentally cut off his pinky. When he was asked what he remebers, he said:*
    
    <br>
    
    #### ***No Long-lived Branches!***
    
    <br>
    
    Down below you will soon learn additional commands for working effectively with
    Git in the command-line. 
    
    These commands are neither complex nor are they used nearly as frequently as the three simple commands we have just worked with when pushing changes up to the remote server. 
    In fact, it should be stressed that it is BEST PRACTICE to commit and push code regularly. 
    
    The motto is: No Long-lived Branches. You likely will only benefit from following it.
    
    In an ideal world, each commit you make should contain an isolated and complete change. 
    
    This helps both you and your teammates out, in both the short- and long-term. 
    
    Keeping your changes isolated and complete allows others to consume and understand
    the changes in a more digestible manner, promoting collaboration and reducing the 
    "noise" along with any confusion or disouragement that might come along with it.
    
    Not only will you be helping out your teammates, by following this best practice you 
    will minimize future frustrations and wasted time. 
    
    Without doubt, you will make mistakes, and those mistakes will be committed up to a remote repository. 
    
    If you have done well to make complete, isolated commits, then finding, targeting, and reverting
    only what is necessary will be easier--and, in many cases, significantly reduce the 
    pain experienced when needing to revert changes, whether in your own remote branch or in the default branch or other shared branches (main, master, develop).
            
    Not only will there be cases where you have committed unwanted code/functionality
    up to some remote, but business requirements change and you are asked to revert some
    recently added feature or functionality.
            
    Frequently pushing code up will help keep your changes backed up as well. 
    While it's necessary and advantageous to coordinate with your team, there are many 
    opportunities for headache around maintaining the team's codebases to occur--even with
    careful planning, there's still the chance two or more people make changes to the same 
    file(s), and run into a mess when getting changes merged into a shared remote branch.
    
    ### Teams at Vizient (branch strat)
    ----
    
    There are many different teams at Vizient, which means things are done in many different ways.
    
    Here are some statistics giving us an indication of which branching strategy teams are using at Vizient. 
    
    While this metric might be helpful, it should be noted that this doesn't fully represent the strategies employed at Vizient. 
    
    stats here
    
    ## Other flow used at Vizient
    
    ## Other flow used at Vizient
    
    ## Github flow -- hybrid/atlassian strategy
    
    ### Team Agreement
    ---
    
    Speak with your team and read  your team's Team Agreement
        
    Every team at Vizient should have a working Team Agreement they have created.
    
    The content of this document is agreed upon by the team, and outlines what is expected of each and every member.
    
    The speciic contents within depend on what the team has agreed upon, but there's 
    likely a section outlining how the engineers on the team collaborate, including the 
    branching strategy, naming and commit message conventions, release cadences, and other important information to promote your success on that team. 
    
    If there's not such a section, or if the team agrees it could be more robust, set up a meeting to discuss and take appropriate action. 
    
    Either way, you should be aware of what conventions and standards they have agreed to follow, and take note. 
    
    For this guide, the goal is to be both comprehensive and promote best practice. 
    
    Given that there are multiple different brancing strategies observed across teams at Vizient, it is beneficial to discuss the most popular ones, along with best practices both general and strategy-specific. 
    
    Although I make an attempt to not be overly opinionated, I would note that any team 
    using the old (~2010) Git flow branching strategy to ask their team if it makes sense to try something different. If this works well for your team, then, great, don't worry about it.
             
    If you are not familiar with your team's document, please reach out to your Product Owner or another personon your team.
    
    ## Don't be afraid to suggest ideas to improve your team's branching strategy
    
    If you have ideas, share them with your team. Try not to be shy or worry too much about how you will be received. It will almost certainly be positive, and it doesn't particularly matter anyway.
    
    Nothing is more valued than ideas, and of course, you need to share the ideas for them to blossom into some real-world difference.
    
    With that said, take time to investigate your team's Team Agreement document, and see if there are any glaring areas needing attention.
    
    ## Considerations with CI/CD and releases
    
    ## Best practices for branch naming
    
    The goal with a branch name is to allow others to learn about ongoing work at a           glance.
    
    
    > git checkout -b feature/867539_refactor-unit-tests
    
    > git checkout spike/862044_try-splitIO
    
    # Convention for commit message format
    
    > #### You will need to check with your team -- try team agreement first. 
    
    ## Explain making changes and VCS principles
    >        
    
    ## Feature Branches
    
    > Having clear, complete, and isolated commits promotes your 
    
    
    
    ### Checking Installation
    
    ---
    
    
    ### Workflow (main)
    
    ---
    
    
    ### Argument for using the command-line
    
    ---
    
    
    ### Other Useful Actions in Git
    
    ---
    
    
    ### Useful but less frequently used commands
    
    ---
    
    shell
    $ git version  
    
    
    shell
    $ git status
    
    
    shell
    $ git fetch
    
    
    shell
    $ git pull
    
    
    shell
    $ git add .
    
    
    shell
    $ git add *
    
    
    shell
    $ git commit -m 'Your commit message here'
    
    
    shell
    $ git push
    
    
    shell
    $ git push --set-remote origin <branch_name>
    
    
    shell
    $ git checkout <branch_name>
    
    
    shell
    $ git checkout -b <branch_name>
    `

	out, err := glamour.Render(in, "dark")
	if err != nil {
		fmt.Println(out)
	}
}
