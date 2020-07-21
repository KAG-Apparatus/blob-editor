using System.Collections.Generic;
using System;
using System.Text.RegularExpressions;
using System.Text;

namespace Blob_Editor
{

    public enum LineType
    {
        Comment,
        EmptyLine,
        EmptyEntry,
        RegularEntry,
        AvulseEntry
    }

    public interface Element
    {
        string Key();
        List<string> Value();
        string Print();
    }

    public class Comment : Element
    {
        private string value;

        public Comment(string value)
        {
            this.value = value;
        }

        public List<string> Value()
        {
            return new List<string>(new string[] { this.value });
        }

        public string Key()
        {
            return null;
        }

        public string Print()
        {
            return this.value;
        }
    }

    public class Empty : Element
    {
        public string Key()
        {
            return null;
        }

        public string Print()
        {
            return "";
        }

        public List<string> Value()
        {
            return null;
        }
    }

    public class Entry : Element
    {
        private string key;
        private List<string> values;

        public Entry(string key, List<string> values)
        {
            this.key = key;
            this.values = values;
        }

        public Entry(string line)
        {
            string pattern = @"(\@*\$*\w+)\ +=\ +(.*)";
            Regex regex = new Regex(pattern, RegexOptions.IgnoreCase);
            Match match = regex.Match(line);
            if (match.Success)
            {
                this.key = match.Groups[1].Captures[0].ToString();
                string value = match.Groups[2].Captures[0].ToString();
                this.values = new List<string>(new string[] { value.Trim() });
            }
            else
            {
                throw new Exception("Invalid entry");

            }
        }

        public void Append(string value)
        {
            this.values.Add(value.Trim());
        }

        public string Key()
        {
            return key;
        }

        public List<string> Value()
        {
            return values;
        }

        public string Print()
        {
            StringBuilder builder = new StringBuilder();
            builder.AppendFormat("{0} = ", this.key);
            builder.AppendJoin(',', this.values);
            return builder.ToString();
        }
    }
}