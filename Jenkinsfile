// define the bitbucket project + repos we want to build
def bitbucket_project = 'Lagrange-Labs'
def bitbucket_repos = ['modmesh']

// create a pipeline job for each of the repos and for each feature branch.
for (bitbucket_repo in bitbucket_repos)
{
  multibranchPipelineJob("${bitbucket_repo}-ci") {
    // configure the branch / PR sources
    branchSources {
      branchSource {
        source {
          bitbucket {
            repoOwner("${bitbucket_project.toUpperCase()}")
            repository("${bitbucket_repo}")
            serverUrl("https://bitbucket.acme.com/")
            traits {
              headWildcardFilter {
                includes("master release/* feature/* bugfix/*")
                excludes("")
              }
            }
          }
        }
        strategy {
          defaultBranchPropertyStrategy {
            props {
              // keep only the last 10 builds
              buildRetentionBranchProperty {
                buildDiscarder {
                  logRotator {
                    daysToKeepStr("-1")
                    numToKeepStr("10")
                    artifactDaysToKeepStr("-1")
                    artifactNumToKeepStr("-1")
                  }
                }
              }
            }
          }
        }
      }
    }
    // discover Branches (workaround due to JENKINS-46202)
    configure {
      def traits = it / sources / data / 'jenkins.branch.BranchSource' / source / traits
      traits << 'com.cloudbees.jenkins.plugins.bitbucket.BranchDiscoveryTrait' {
        strategyId(3) // detect all branches
      }
    }

    // check every minute for scm changes as well as new / deleted branches
    triggers {
      periodic(1)
    }
    // don't keep build jobs for deleted branches
    orphanedItemStrategy {
      discardOldItems {
        numToKeep(-1)
      }
    }
  }
}