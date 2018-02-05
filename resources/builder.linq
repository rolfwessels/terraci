<Query Kind="Program" />

void Main()
{
	var folder = @"D:\Work\Home\Go\src\github.com\rolfwessels\continues-terraforming\resources\sample";
	var files = Directory.GetFiles(folder, "*.*", SearchOption.AllDirectories);
	var filesToRemove = files.Where(x => x.EndsWith(".tf") && !x.EndsWith("main.tf") && !x.EndsWith("output.tf") && !x.EndsWith("variables.tf")).Dump("remove").Select(x => { File.Delete(x); return true; }).ToArray();
	//var filesToRemoveStateFiles = files.Where(x => x.Contains(".tfstate") ).Dump("remove").Select(x => {File.Delete(x);  return true; }).ToArray();
	
	var foldersToMod = files.Where(x => x.EndsWith(".tf")).Where(x=> !x.Contains("global") && !x.Contains("modules")).Select(x => Path.GetDirectoryName(x)).Distinct().Dump().ToArray();
	var mainTemplate = File.ReadAllText(files.FirstOrDefault(x => x.EndsWith(@"global\main.tf")));
	var outputTemplate = File.ReadAllText(files.FirstOrDefault(x => x.EndsWith(@"global\output.tf")));
	var variablesTemplate = File.ReadAllText(files.FirstOrDefault(x => x.EndsWith(@"global\variables.tf")));
	foreach (var toFolder in foldersToMod)
	{
		var varContent = variablesTemplate + "variable \"region\" { }\n";
		if (toFolder.Contains("env"))
		{
			varContent = varContent + "variable \"env\" { }\n";
		}
		File.WriteAllText(Path.Combine(toFolder, "variables.tf"), varContent);
		var name = "terraform-${var.region}-"+Path.GetFileNameWithoutExtension(toFolder);
		var groupName = "security_"+Path.GetFileNameWithoutExtension(toFolder);
		if (toFolder.Contains("env"))
		{
			name = "terraform-${var.region}-${var.env}-"+Path.GetFileNameWithoutExtension(toFolder);
		}
		File.WriteAllText(Path.Combine(toFolder, "main.tf"), mainTemplate.Replace("terraform-global", name).Replace("default", groupName));
		File.WriteAllText(Path.Combine(toFolder, "output.tf"), outputTemplate.Replace("terraform-global", name).Replace("default", groupName).Replace("\"name\"", $"\"{groupName}_name\""));
	}
}
// variable "aws_region" { }
// variable "env" { }
//

