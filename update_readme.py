import sys
import os
import re

challenges = None
directory = 'challenges'

def getListOfChallenges(dirName):

    listOfChallenges = os.listdir(dirName)
    allChallenges = list()
    for entry in listOfChallenges:
        fullPath = os.path.join(dirName, entry)
        if os.path.isdir(fullPath):
            allChallenges = allChallenges + getListOfChallenges(fullPath)
        else:
            if ".md" not in entry:
                allChallenges.append(fullPath)        
                
    return allChallenges

def getChallengesCount():
    
    global challenges

    if challenges == None:
        challenges = getListOfChallenges(directory)

    return len(challenges)

def updateReadme():

    readmeFile = 'README.md'
    with open(readmeFile, 'r') as file:
        filedata = file.read()    

    total_challenges = getChallengesCount()
    filedata = re.sub('\*\*Completed Challenges:\*\*.*', '**Completed Challenges:** {}'.format(total_challenges),filedata)

    with open(readmeFile, 'w') as file:
        file.write(filedata)

if __name__ == "__main__":
    updateReadme()