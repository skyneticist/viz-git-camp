## Git Guide


# The core function of a version control system is to commit/share changes made on your working copy to a shared server.

# With git, it takes at least three separate commands to share a commit with your colleagues:

	- `git add .` or `git add *`

	- `git commit -m "message to associate with commit"`

	- `git push`

	
    The first step adds all files with changes from the working copy of code on your machine      to the `staging area`. 

    From here, these changes can be committed to your local repository. We do this in the         second step, passing `-m` command-line flag to specify we are providing a message             describing the changes of the commit, inline, with this commit command.
 
    Finally we simply run `git push` to push the changes on our local repository up to the        remote repository for others on your team to pull down on their side. 

    Notice that we could have used `*` or `.` with the first command. Each command results in     different behaviors and we should consider the differences now. 
    
    In a nutshell, `git add .` adds all files, folders, and subfolders, including `.gitignore`    and all other hidden or `dot` files. `git add *` does virtually the same thing, with the 	one difference being it will not add any dot files (files beginning with a dot `.` -- 
    hidden files.

    You likely won't run into many cases where you wish to ignore changes made to hidden or       dot files, but using `git add *` would be the command to push any other file changes.

    With that said, you will principally use `git add .` to add local file changes to the 
    staging area. Although not particularly imporatant, it should be noted that `*` is
    expanded for you, while `.` is a git-specific argument that it uses to add all changes.

# Branching Strategy/Model

	Down below you will soon learn additional commands for working effectively with
	Git in the command-line. These commands are neither complex nor are they used 
        nearly as frequently as the three simple commands we have just worked with when 
	pushing changes up to the remote server. In fact, it should be stressed that 
        it is BEST PRACTICE to commit and push code regularly. The motto is "no long-lived
        branches". In an ideal world, each commit you make should contain an isolated and
        complete change. This helps both you and your teammates out, in the short- and long-
        term. 

	Keeping your changes isolated and complete allows others to consume and understand
        the changes in a more digestible manner, promoting collaboration and reducing the 
        "noise" along with any confusion or disouragement that might come along with it. 

        Not only will you be helping out your teammates, by following this best practice you 
        will minimize future frustrations and wasted time. Without doubt, you will make 
        mistakes, and those mistakes will be committed up to a remote repository. If you have
        done well to make complete, isolated commits, then finding, targeting, and reverting
        only what is necessary will be easier--and, in many cases, significantly reduce the 
        pain experienced when needing to revert changes, whether in your own remote branch or         in the default branch or other shared branches (`main`, `master`, `develop`).
        
        Not only will there be cases where you have committed unwanted code/functionality
	up to some remote, but business requirements change and you are asked to revert some
        recently added feature or functionality. 
        
        Frequently pushing code up will help keep your changes backed up as well. 
        While it's necessary and advantageous to coordinate with your team, there are many 
        opportunities for headache around maintaining the team's codebases to occur--even with
        careful planning, there's still the chance two or more people make changes to same 
        file(s), and run into a mess when getting changes merged into a shared remote branch.

        

## Teams at Vizient (branch strat)

There are many different teams at Viziet, which means things are done in many different ways.

Here are some statistics giving us an indication of which branching strategy teams are using 
at Vizient. While this metric is helpful, it should be noted that this doesn't fully representstrategies employed at Vizient. 

```stats here```

## Other flow used at Vizient

## Other flow used at Vizient

## Github flow -- hybrid/atlassian strategy

## Speak with team and read team agreement
	
	Every team at Vizient should have a working Team Agreement they have created.
        The content of this document is agreed upon by the team, and outlines what is expected        of each and every member. 

	The speciic contents within depend on what the team has agreed upon, but there's 
        likely a section outlining how the engineers on the team collaborate, including the 
        branching strategy, naming and commit message conventions, release cadences, and other        important information to promote your success on that team. 

	If there's not such a section, or if the team agrees it could be more robust, set up a        meeting to discuss and take appropriate action. 

	Either way, you should be aware of what conventions and standards they have agreed to 	      follow, and take note. 

        For this guide, the goal is to be both comprehensive and promote best practice. 
        Given that there are multiple different brancing strategies observed across teams at
        Vizient, it is beneficial to discuss the most popular ones, along with best practices         both general and strategy-specific. 

        Although I make an attempt to not be overly opinionated, I would note that any team 
        using the old (~2010) Git flow branching strategy to ask their team if it makes sense
        to try something different. If this works well for your team, then, great, don't worry        about it. 
         

	If you are not familiar with your team's document, please reach out to your Product           Owner or another person on your team.

        

## Don't be afraid to suggest ideas to improve your team's branching strategy

## Considerations with CI/CD and releases

## Best practices for branch naming

	The goal with branch names is to allow others to learn about ongoing work at a                glance.

	`git checkout -b feature/867539_refactor-unit-tests`
	`git checkout spike/862044_try-splitIO`

# Convention for commit message format

	You will need to check with your team -- try team agreement first. 


## Explain making changes and VCS principles

## 
        
## Feature Branches



        
        Having clear, complete, and isolated commits promotes your   
