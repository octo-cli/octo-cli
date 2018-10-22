Service Issues {
  Command Edit {
    ArgNames = ["Owner", "Repo", "Number"]
  }
  Command Create {}
  Command Lock {}
  Command AddLabelsToIssue {
    RoutesName = "add-labels"
    ArgNames = ["Owner", "Repo", "Number", "Labels"]
  }
  Command List{
    ArgNames = ["All"]
  }
}

Service Organizations {
  RouteName = "orgs"
  Command Get {}
}

Service Repositories {
  RouteName = "repos"
  Command Get {}
}
