import sys
import os

def getListOfFiles(dirName):

    listOfFile = os.listdir(dirName)
    allFiles = list()
    for entry in listOfFile:
        fullPath = os.path.join(dirName, entry)
        if os.path.isdir(fullPath):
            allFiles = allFiles + getListOfFiles(fullPath)
        else:
            allFiles.append(fullPath)        
                
    return allFiles

if __name__ == "__main__":
    files = getListOfFiles('challenges')

    for file in files:
        print(file)