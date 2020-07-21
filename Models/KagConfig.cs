using System.Collections.Generic;

namespace Blob_Editor
{
    public class KagConfig
    {
        private List<Element> elements;

        public KagConfig(List<Element> elements)
        {
            this.Elements = elements;
        }

        public List<Element> Elements { get => elements; set => elements = value; }
    }
}